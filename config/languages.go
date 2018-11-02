package config

// Language represents programming languages to generate code for
type Language string

// Languages is a list of all available languages
var Languages = []Language{Go}

const (
	// Go represents the programming language Go
	// https://golang.org/
	// https://github.com/golang/protobuf
	Go = Language("go")
)
