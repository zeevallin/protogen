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
	KWSource   = "source"
	KWLanguage = "language"
	KWGenerate = "generate"
)

var (
	// Keywords represent all valid keywords
	Keywords = map[string]Type{
		"source":   SOURCE,
		"language": LANGUAGE,
		"generate": GENERATE,
	}
)

// LookupIdentifier returns the type of the identifier
func LookupIdentifier(id string) Type {
	if kw, ok := Keywords[id]; ok {
		return kw
	}
	return IDENTIFIER
}
