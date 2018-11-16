package source

// RefType is a type of reference
type RefType int

const (
	// Branch defines a branch type
	Branch = RefType(iota)

	// Version is a version tag
	Version
)

// Ref defines the source variant
type Ref struct {
	Type RefType
	Name string
}
