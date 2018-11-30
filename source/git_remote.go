package source

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"path"
	"strings"

	billy "gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/plumbing/format/config"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/storage"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

const (
	gitSourceErrFmt        = "cannot create new git source: %v"
	gitInitErrFmt          = "cannot initialize git source: %v"
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
		return nil, fmt.Errorf(gitSourceErrFmt, err)
	}
	switch url.Scheme {
	case "":
		url.Scheme = defaultScheme
	}

	p := path.Join(os.TempDir(), "protogen")
	treePath := path.Join(p, "tree")
	repoPath := path.Join(p, "repo")

	storer := filesystem.NewStorage(osfs.New(treePath), cache.NewObjectLRUDefault())

	return &RemoteGitSource{
		URL:      url.String(),
		fs:       osfs.New(repoPath),
		path:     p,
		repoPath: repoPath,
		treePath: treePath,
		storer:   storer,
		logger:   logger,
	}, nil
}

// RemoteGitSource represents a remote git source tree
type RemoteGitSource struct {
	URL string

	path     string
	repoPath string
	treePath string

	fs     billy.Filesystem
	storer storage.Storer
	repo   *git.Repository
	wt     *git.Worktree

	logger *log.Logger
}

// Init initialises the git repository
func (rgs *RemoteGitSource) Init() (err error) {
	rgs.logger.Println("initialising remote git tree")
	rgs.logger.Printf("removing everything at: %s\n", rgs.path)
	err = os.RemoveAll(rgs.path)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Printf("creating directory for tree: %s\n", rgs.treePath)
	err = os.MkdirAll(rgs.treePath, 0700)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Printf("creating directory for repo: %s\n", rgs.repoPath)
	err = os.MkdirAll(rgs.repoPath, 0700)
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Println("setting up git configuration")
	err = rgs.setupGitconfig()
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Println("cloning git repository")
	rgs.repo, err = git.Clone(rgs.storer, rgs.fs, &git.CloneOptions{
		URL:  rgs.URL,
		Tags: git.AllTags,
	})
	if err != nil {
		return fmt.Errorf(gitInitErrFmt, err)
	}
	rgs.logger.Println("retrieving work tree for git repo")
	rgs.wt, err = rgs.repo.Worktree()
	return
}

// Checkout checks out a specific hash
func (rgs *RemoteGitSource) Checkout(hash string) error {
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

// Packages lists the packages in the repository
func (rgs *RemoteGitSource) Packages() ([]string, error) {
	return nil, nil
}

// PackageVersions lists the versions for a given package
func (rgs *RemoteGitSource) PackageVersions(pkg string) ([]string, error) {
	return nil, nil
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

func (rgs *RemoteGitSource) setupGitconfig() (err error) {
	usr, err := user.Current()
	if err != nil {
		return
	}

	path := path.Join(usr.HomeDir, ".gitconfig")
	f, err := os.Open(path)
	if err != nil {
		return
	}

	dec := config.NewDecoder(f)
	gitconfig := &config.Config{}
	err = dec.Decode(gitconfig)
	if err != nil {
		return
	}

	cfg, err := rgs.storer.Config()
	if err != nil {
		return
	}

	cfg.Raw = gitconfig
	err = rgs.storer.SetConfig(cfg)
	return
}
