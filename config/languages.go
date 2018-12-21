package config

import (
	"github.com/zeeraw/protogen/config/go"
	"github.com/zeeraw/protogen/config/swift"
)

// Language represents programming languages to generate code for
type Language string

// Languages is a list of all available languages
var Languages = []Language{
	Go,
	Swift,
}

const (
	// Go represents the programming language Go
	// https://golang.org/
	// https://github.com/golang/protobuf
	Go = Language("go")

	// Swift represents the programming language Swift
	// https://swift.org/
	// https://github.com/apple/swift-protobuf
	Swift = Language("swift")
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
	case Swift:
		return swift.Config{}
	}
	panic("cannot find language")
}
