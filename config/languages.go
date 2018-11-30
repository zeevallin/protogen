package config

import (
	"github.com/zeeraw/protogen/config/go"
)

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

// ForLanguage returns the appropriate configuration for a given programming language
func ForLanguage(lang Language) interface{} {
	switch lang {
	case Go:
		return golang.Config{
			Plugins: []golang.Plugin{
				golang.GRPC,
			},
		}
	}
	panic("cannot find language")
}
