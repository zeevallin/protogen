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

// Init is a no-op
func (s *MockSource) Init() error {
	return nil
}

// PathTo returns the path of the mock fixture package
func (s *MockSource) PathTo(pkg string) string {
	return path.Join(s.RootPath(), pkg)
}

// RootPath returns the root path of the mock fixtures
func (s *MockSource) RootPath() string {
	return s.repoPath
}

// Checkout is a no-op
func (s *MockSource) Checkout(hash string) error {
	return nil
}

// HashForRef is a no-op that returns an empty string
func (s *MockSource) HashForRef(ref Ref) (string, error) {
	return "", nil
}
