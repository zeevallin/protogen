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
)

var (
	// Keywords represent all valid keywords
	Keywords = map[string]Type{
		KWSource:   SOURCE,
		KWLanguage: LANGUAGE,
		KWGenerate: GENERATE,
		KWOutput:   OUTPUT,
	}
)

// LookupIdentifier returns the type of the identifier
func LookupIdentifier(id string) Type {
	if kw, ok := Keywords[id]; ok {
		return kw
	}
	return IDENTIFIER
}
