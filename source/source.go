package source

import (
	"os"
)

func init() {
	WorkDir := os.Getenv("PROTOGEN_WORKDIR")
	if WorkDir == "" {
		WorkDir = "/usr/local/var/protogen"
	}
}

var (
	// WorkDir is the dictionary in which we operate
	WorkDir string
)

// Source defines behaviour of the interaction with a proto repository
type Source interface {
	Init() error
	PathTo(pkg string) string

	Checkout(hash string) error
	HashForRef(ref Ref) (string, error)

	Packages() ([]string, error)
	PackageVersions(pkg string) ([]string, error)
}
