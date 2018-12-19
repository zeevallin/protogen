package source

var (
	// WorkDir is the dictionary in which we operate
	WorkDir string
)

// Source defines behaviour of the interaction with a proto repository
type Source interface {
	Init() error
	PathTo(pkg string) string
	RootPath() string

	Checkout(hash string) error
	HashForRef(ref Ref) (string, error)
}
