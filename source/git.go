package source

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

const (
	urlSchemeDefault       = "https"
	fmtErrGitHashForTag    = "cannot get commit hash for package tag %s: %v"
	fmtErrGitHashForBranch = "could not resolve hash for branch: %v"
	remoteOriginPrefix     = "refs/remotes/origin/"
)

var (
	// ErrURLParse happens when failing to parse the remote as a url
	ErrURLParse = errors.New("error parsing remote git source url")

	// ErrOpenRepo happens when you try to open a repo on disk and it doesnt exist
	ErrOpenRepo = errors.New("error opening repository on disk")

	// ErrCloneRemote happens when you try to clone a repo and something fails
	ErrCloneRemote = errors.New("error cloning repository")

	// ErrFetchRemote happens when trying to fetch the remote
	ErrFetchRemote = errors.New("error fetching the remote")

	// ErrMakeWorkTreeTmpDir happens when trying to create the work tree directory for the temp work tree
	ErrMakeWorkTreeTmpDir = errors.New("error creating the work tree directory")

	// ErrCleanWorkTreeTmpDir happens when you can't clean the work tree from disk
	ErrCleanWorkTreeTmpDir = errors.New("error removing the work tree directory")

	// ErrMakeSrcTree happens when you fail to create the directory for the remote source tree cache
	ErrMakeSrcTree = errors.New("error creating the remote source tree directory")

	// ErrOpenWorkTree happens when trying to open the git work tree
	ErrOpenWorkTree = errors.New("error opening the git work tree")

	// ErrUnknownRefType happens when the ref type is unknown
	ErrUnknownRefType = errors.New("ref type unknown")

	// ErrFindingBranch happens when the branch cannot be found
	ErrFindingBranch = errors.New("error finding branch")

	// ErrCheckLocalRepo happens when the local repository cannot be resolved
	ErrCheckLocalRepo = errors.New("error checking local repository")

	// TmpWorkDir represents the directory used to checkout the work tree
	TmpWorkDir = path.Join(os.TempDir(), "protogen")
)

// NewMockGitSource returns a local mock source
func NewMockGitSource(p string) *GitSource {
	return &GitSource{
		Remote: nil,
		Repo:   p,
		Tree:   "",
	}
}

// NewLocalGitSource initialises a new git source
func NewLocalGitSource(p string) (*GitSource, error) {
	tree, err := getLocalRepo(path.Join(p, ".git"))
	if err != nil {
		return nil, err
	}

	return &GitSource{
		Remote: nil,
		Repo:   TmpWorkDir,
		Tree:   tree,
	}, nil
}

func getLocalRepo(p string) (string, error) {
	log.Printf("checking local git repo: %s\n", p)
	fi, err := os.Stat(p)
	if err != nil {
		return "", ErrCheckLocalRepo
	}
	if !fi.IsDir() {
		log.Println("repo tree is not directory, following trail")
		f, err := ioutil.ReadFile(p)
		if err != nil {
			return "", ErrCheckLocalRepo
		}
		trail := strings.TrimSpace(string(f))
		trail = strings.TrimPrefix(trail, "gitdir: ")
		root := strings.TrimRight(p, "/.git")
		trail, err = filepath.Abs(path.Join(root, trail))
		if err != nil {
			return "", ErrCheckLocalRepo
		}
		log.Printf("found tree trail: %s\n", trail)
		return getLocalRepo(trail)
	}
	return p, nil
}

// NewRemoteGitSource initialises a new git source
func NewRemoteGitSource(remote string) (*GitSource, error) {
	u, err := url.Parse(remote)
	if err != nil {
		return nil, ErrURLParse
	}
	switch u.Scheme {
	case "":
		u.Scheme = urlSchemeDefault
	}
	tree := path.Join(WorkDir, "src", u.Path)
	if err := os.MkdirAll(tree, 0700); err != nil {
		return nil, ErrMakeSrcTree
	}

	return &GitSource{
		Remote: u,
		Repo:   TmpWorkDir,
		Tree:   tree,
	}, nil
}

// GitSource represents a git source
type GitSource struct {
	// Remote represents the git remote
	Remote *url.URL

	// Repo represents the location of the current repository content
	Repo string

	// Three represents the git tree, normally the .git directory
	Tree string
}

// PathTo returns the path to a package on the source
func (gs *GitSource) PathTo(pkg string) string {
	return path.Join(gs.Root(), pkg)
}

// Root returns the root path for importing dependent packages
func (gs *GitSource) Root() string {
	return gs.Repo
}

// InitRepo initialises a git repo
func (gs *GitSource) InitRepo() (Repo, error) {
	repoFS := osfs.New(gs.Repo)
	treeFS := osfs.New(gs.Tree)
	storer := filesystem.NewStorage(treeFS, cache.NewObjectLRUDefault())
	log.Printf("opening local git repo: %s\n", gs.Tree)
	repo, err := git.Open(storer, repoFS)
	if err != nil && gs.Remote == nil {
		return nil, ErrOpenRepo
	}
	if err != nil && gs.Remote != nil {
		log.Printf("cloning remote git repo: %s\n", gs.Remote.String())
		repo, err = git.Clone(storer, repoFS, &git.CloneOptions{
			URL:  gs.Remote.String(),
			Tags: git.AllTags,
		})
		if err != nil {
			return nil, ErrCloneRemote
		}
	}
	if err == nil && gs.Remote != nil {
		log.Printf("fetching remote git repo: %s\n", gs.Remote.String())
		err := repo.Fetch(&git.FetchOptions{
			Force: true,
			Tags:  git.AllTags,
		})
		switch err.(error) {
		case git.NoErrAlreadyUpToDate:
		default:
			return nil, ErrFetchRemote
		}
	}
	log.Printf("opening work tree: %s\n", gs.Repo)
	wt, err := repo.Worktree()
	if err != nil {
		return nil, ErrOpenWorkTree
	}
	return &GitRepo{
		storer: storer,
		path:   gs.Repo,
		repo:   repo,
		wt:     wt,
	}, nil
}

// GitRepo represents a worktree
type GitRepo struct {
	path   string
	storer *filesystem.Storage
	repo   *git.Repository
	wt     *git.Worktree
}

// Clean cleans the repo work tree from disk
func (gr *GitRepo) Clean() error {
	log.Printf("removing directory for repo: %s\n", gr.path)
	err := os.RemoveAll(gr.path)
	if err != nil {
		return ErrCleanWorkTreeTmpDir
	}
	return nil
}

// Checkout will checkout the repository
func (gr *GitRepo) Checkout(ref Ref) error {
	log.Printf("creating directory for repo: %s\n", gr.path)
	if err := os.MkdirAll(TmpWorkDir, 0700); err != nil {
		return ErrMakeWorkTreeTmpDir
	}
	hash, err := gr.hashForRef(ref)
	if err != nil {
		return err
	}
	log.Printf("checking out hash in repo: %s\n", hash)
	return gr.wt.Checkout(&git.CheckoutOptions{
		Create: false,
		Force:  true,
		Hash:   plumbing.NewHash(hash),
	})
}

func (gr *GitRepo) hashForRef(ref Ref) (string, error) {
	log.Printf("retrieving hash for ref: %s (%v)\n", ref.Name, ref.Type)
	switch ref.Type {
	case Version:
		return gr.hashForTag(ref.Name)
	case Branch:
		return gr.hashForBranch(ref.Name)
	default:
		return "", ErrUnknownRefType
	}
}

// hashForBranch derives the commit hash from a branch on the origin
func (gr *GitRepo) hashForBranch(branch string) (string, error) {
	refs, err := gr.storer.IterReferences()
	if err != nil {
		return "", fmt.Errorf(fmtErrGitHashForBranch, err)
	}
	it := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs)
	var hash string
	err = it.ForEach(func(ref *plumbing.Reference) error {
		b := strings.TrimPrefix(ref.Name().String(), remoteOriginPrefix)

		if b == branch {
			hash = ref.Hash().String()
			return nil
		}
		return ErrFindingBranch
	})
	if err != nil {
		return "", fmt.Errorf(fmtErrGitHashForBranch, err)
	}
	return hash, nil
}

// hashForTag derives the commit hash from a tag for a specific package
func (gr *GitRepo) hashForTag(tag string) (string, error) {
	ref, err := gr.repo.Tag(tag)
	if err != nil {
		return "", fmt.Errorf(fmtErrGitHashForTag, tag, err)
	}
	return ref.Hash().String(), nil
}

// MockRepo represents a no-op collection of methods to satisfy the repo interface
type MockRepo struct{}

// Clean is a no-op
func (s *MockRepo) Clean() error {
	return nil
}

// Checkout is a no-op
func (s *MockRepo) Checkout(ref Ref) error {
	return nil
}
