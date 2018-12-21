package source

var (
	// WorkDir is the dictionary in which we operate
	WorkDir string
)

// Source defines behaviour for a source
type Source interface {
	Root() string
	PathTo(pkg string) string
	InitRepo() (Repo, error)
}

// Repo defines behaviour for a git repository
type Repo interface {
	Clean() error
	Checkout(ref Ref) error
}
