package token

// Type defines the type of token as a string
type Type string

// Token defines a token in the file with its type and literal value
type Token struct {
	Type    Type
	Literal string
}

const (
	// General

	// ILLEGAL is an unknown token
	ILLEGAL = "ILLEGAL"

	// EOF defines the end of the file
	EOF = "EOF"

	// CONFIGURATION defines the token for a top level file
	CONFIGURATION = "CONFIGURATION"

	// Keywords

	// SOURCE defines the keyword for the protogen remote source
	SOURCE = "SOURCE"

	// LANGUAGE defines the keyword for the output language
	LANGUAGE = "LANGUAGE"

	// GENERATE defines the keywod for what package to generate
	GENERATE = "GENERATE"

	// OUTPUT defines the keywod for where to generate packages to
	OUTPUT = "OUTPUT"

	// PLUGIN defines the keyword for adding a plugin for go code generation
	PLUGIN = "PLUGIN"

	// PATH defines the keyword for defining the import pathing for go code generation
	PATH = "PATH"

	// OPTION defines the keyword for specifying options for swift code generation
	OPTION = "OPTION"

	// Literals

	// VERSION defines a version value
	VERSION = "VERSION"

	// IDENTIFIER defines a package path
	IDENTIFIER = "IDENTIFIER"

	// Delimiters

	// WHITESPACE defines regular whitespace between tokens
	WHITESPACE = "WHITESPACE"

	// NEWLINE defines a new line in the file
	NEWLINE = "NEWLINE"

	// LEFTBRACE defines a left brace "{"
	LEFTBRACE = "LEFTBRACE"

	// RIGHTBRACE defines a right brace "}"
	RIGHTBRACE = "RIGHTBRACE"
)

const (
	// KWSource is the explicit keyword for source
	KWSource = "source"

	// KWLanguage is the explicit keyword for language
	KWLanguage = "language"

	// KWGenerate is the explicit keyword for generate
	KWGenerate = "generate"

	// KWOutput is the explicit keyword for output
	KWOutput = "output"

	// KWPlugin is the explicit keyword for a go plugin
	KWPlugin = "plugin"

	// KWPath is the explicit keyword for a go import path
	KWPath = "path"

	// KWOption is the explicit keyword for a language option
	KWOption = "option"
)

var (
	// Keywords represent all valid keywords
	Keywords = map[string]Type{
		KWSource:   SOURCE,
		KWLanguage: LANGUAGE,
		KWGenerate: GENERATE,
		KWOutput:   OUTPUT,
		KWPath:     PATH,
		KWPlugin:   PLUGIN,
		KWOption:   OPTION,
	}
)

// LookupIdentifier returns the type of the identifier
func LookupIdentifier(id string) Type {
	if kw, ok := Keywords[id]; ok {
		return kw
	}
	return IDENTIFIER
}
