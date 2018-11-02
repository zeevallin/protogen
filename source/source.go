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
	Path() string

	Update() error
	Checkout(tag string) error
}
