package source

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type GitSource struct {
	path string
}

type RemoteGitSource struct {
	GitSource
	URL string

	repository *git.Repository
}

// Update the remote git repository
func (gs *RemoteGitSource) Update() error {
	var err error
	if gs.repository != nil {
		err := gs.repository.Fetch(&git.FetchOptions{
			Force: true,
			Tags:  git.AllTags,
		})
		return err
	}
	gs.repository, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: gs.URL,
	})
	return err
}

// Checkout checks out a specific tag or branch
func (gs *RemoteGitSource) Checkout(tag string) error {
	return nil
}

// Path returns the path to the repository on disk
func (gs *RemoteGitSource) Path() string {
	return ""
}

type LocalGitSource struct {
	GitSource
}

// Update does nothing since the local source is
func (gs *LocalGitSource) Update() error {
	// noop
	return nil
}
