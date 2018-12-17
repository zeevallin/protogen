package source

import (
	"log"
	"path"
)

// MockSource represents a mocked source on disk
type MockSource struct {
	repoPath string
}

// NewMockSource returns a local mock source
func NewMockSource(logger *log.Logger, p string) (*MockSource, error) {
	return &MockSource{
		repoPath: p,
	}, nil
}

func (s *MockSource) Init() error {
	return nil
}

func (s *MockSource) PathTo(pkg string) string {
	return path.Join(s.RootPath(), pkg)
}

func (s *MockSource) RootPath() string {
	return s.repoPath
}

func (s *MockSource) Checkout(hash string) error {
	return nil
}

func (s *MockSource) HashForRef(ref Ref) (string, error) {
	return "", nil
}

func (s *MockSource) Packages() ([]string, error) {
	return []string{}, nil
}

func (s *MockSource) PackageVersions(pkg string) ([]string, error) {
	return []string{}, nil
}
