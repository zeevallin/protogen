package source

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	billy "gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/storage"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

const (
	gitSourceErrFmt        = "cannot create new git source: %v"
	gitInitErrFmt          = "cannot initialise git source: %v"
	gitHashForTag          = "cannot get commit hash for package tag %s: %v"
	gitHashForBranchErrFmt = "could not resolve hash for branch: %v"

	defaultScheme = "https"
	originPrefix  = "refs/remotes/origin/"
)

// NewRemoteGitSource returns a remote git source manager
func NewRemoteGitSource(logger *log.Logger, u string) (*RemoteGitSource, error) {
	logger.Printf("parsing remote git source: %s\n", u)
	url, err := url.Parse(u)
	if err != nil {
		logger.Printf("error parsing url: %s\n", err)
		return nil, fmt.Errorf(gitSourceErrFmt, err)
	}
	switch url.Scheme {
	case "":
		url.Scheme = defaultScheme
	}

	repoPath := path.Join(os.TempDir(), "protogen")
	repoFs := osfs.New(repoPath)

	treePath := path.Join(WorkDir, "src", url.Path)
	treeFs := osfs.New(treePath)

	storer := filesystem.NewStorage(treeFs, cache.NewObjectLRUDefault())

	return &RemoteGitSource{
		URL:      url.String(),
		repoFs:   repoFs,
		repoPath: repoPath,
		treePath: treePath,
		treeFs:   treeFs,
		storer:   storer,
		lock:     &sync.Mutex{},
		logger:   logger,
	}, nil
}

// RemoteGitSource represents a remote git source tree
type RemoteGitSource struct {
	URL string

	repoPath string
	treePath string

	repoFs billy.Filesystem
	treeFs billy.Filesystem

	storer storage.Storer
	repo   *git.Repository
	wt     *git.Worktree

	cloned bool
	lock   *sync.Mutex
	logger *log.Logger
}

// Init initialises the git repository
func (rgs *RemoteGitSource) Init() error {
	rgs.lock.Lock()
	defer rgs.lock.Unlock()
	rgs.logger.Println("initialising remote git source")
	var err error
	if !rgs.cloned {
		rgs.logger.Println("remote git source has not been cloned yet")
		err = rgs.clone()
	} else {
		rgs.logger.Println("remote git source is already cloned")
		err = rgs.fetch()
	}
	return err
}

func (rgs *RemoteGitSource) fetch() error {
	rgs.logger.Println("fetching remote git source")
	err := rgs.repo.Fetch(&git.FetchOptions{
		Force: true,
		Tags:  git.AllTags,
	})
	switch err.(error) {
	case git.NoErrAlreadyUpToDate:
		return nil
	}
	return fmt.Errorf(gitInitErrFmt, err)
}

func (rgs *RemoteGitSource) reset() error {
	rgs.logger.Printf("removing directory for repo: %s\n", rgs.repoPath)
	err := os.RemoveAll(rgs.repoPath)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Printf("creating directory for repo: %s\n", rgs.repoPath)
	err = os.MkdirAll(rgs.repoPath, 0700)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	return nil
}

func (rgs *RemoteGitSource) open() (err error) {
	rgs.logger.Println("opening existing git source")
	rgs.repo, err = git.Open(rgs.storer, rgs.repoFs)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	return err
}

func (rgs *RemoteGitSource) clone() error {
	if err := rgs.reset(); err != nil {
		return err
	}
	rgs.logger.Printf("creating directory for tree: %s\n", rgs.treePath)
	err := os.MkdirAll(rgs.treePath, 0700)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Println("cloning git repository")
	rgs.repo, err = git.Clone(rgs.storer, rgs.repoFs, &git.CloneOptions{
		URL:  rgs.URL,
		Tags: git.AllTags,
	})
	if err != nil {
		switch err {
		case git.ErrRepositoryAlreadyExists:
			rgs.logger.Println("repository already exists")
			if err := rgs.open(); err != nil {
				return err
			}
			if err := rgs.fetch(); err != nil {
				return err
			}
		default:
			rgs.logger.Printf("could not clone repository: %v\n", err)
			return fmt.Errorf(gitInitErrFmt, err)
		}
	}
	rgs.cloned = true
	rgs.logger.Println("retrieving work tree for git repo")
	rgs.wt, err = rgs.repo.Worktree()
	return err
}

// Checkout checks out a specific hash
func (rgs *RemoteGitSource) Checkout(hash string) error {
	rgs.lock.Lock()
	defer rgs.lock.Unlock()

	rgs.logger.Printf("checking out hash: %s\n", hash)
	return rgs.wt.Checkout(&git.CheckoutOptions{
		Create: false,
		Force:  true,
		Hash:   plumbing.NewHash(hash),
	})
}

// PathTo returns the path to a package on disk
func (rgs *RemoteGitSource) PathTo(pkg string) string {
	return path.Join(rgs.RootPath(), pkg)
}

// RootPath returns the root path for importing dependent packages
func (rgs *RemoteGitSource) RootPath() string {
	return rgs.repoPath
}

// HashForRef derives the
func (rgs *RemoteGitSource) HashForRef(ref Ref) (string, error) {
	rgs.lock.Lock()
	defer rgs.lock.Unlock()

	rgs.logger.Printf("retrieving hash for ref: %s (%v)\n", ref.Name, ref.Type)
	switch ref.Type {
	case Version:
		return rgs.hashForTag(ref.Name)
	case Branch:
		return rgs.hashForBranch(ref.Name)
	default:
		return "", fmt.Errorf("unknown ref type")
	}
}

// hashForBranch derives the commit hash from a branch on the origin
func (rgs *RemoteGitSource) hashForBranch(branch string) (string, error) {
	refs, err := rgs.storer.IterReferences()
	if err != nil {
		return "", fmt.Errorf(gitHashForBranchErrFmt, err)
	}

	it := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs)

	var hash string
	err = it.ForEach(func(ref *plumbing.Reference) error {
		b := strings.TrimPrefix(ref.Name().String(), originPrefix)

		if b == branch {
			hash = ref.Hash().String()
			return nil
		}
		return errors.New("couldn't find branch")
	})
	if err != nil {
		return "", fmt.Errorf(gitHashForBranchErrFmt, err)
	}

	return hash, nil
}

// hashForTag derives the commit hash from a tag for a specific package
func (rgs *RemoteGitSource) hashForTag(tag string) (string, error) {
	ref, err := rgs.repo.Tag(tag)
	if err != nil {
		return "", fmt.Errorf(gitHashForTag, tag, err)
	}
	return ref.Hash().String(), nil
}
