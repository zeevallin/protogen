package token

const (
	// GOPLUGIN defines the keyword for adding a plugin for go code generation
	GOPLUGIN = "GOPLUGIN"

	// GOPATH defines the keyword for defining the import pathing for go code generation
	GOPATH = "GOPATH"
)

const (
	// KWGoPlugin is the explicit keyword for a go plugin
	KWGoPlugin = "plugin"

	// KWGoPath is the explicit keyword for a go import path
	KWGoPath = "path"
)

var (
	// GoKeywords represent all valid go keywords
	GoKeywords = map[string]Type{
		KWGoPlugin: GOPLUGIN,
		KWGoPath:   GOPATH,
	}
)

// LookupGoIdentifier returns the type of the go identifier
func LookupGoIdentifier(id string) Type {
	if kw, ok := GoKeywords[id]; ok {
		return kw
	}
	return IDENTIFIER
}
