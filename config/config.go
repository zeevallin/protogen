package config

import (
	"github.com/zeeraw/protogen/source"
)

// Config defines the configuration for code generation
type Config struct {
	Packages []*Package
}

// Package defines a package to generate code for
type Package struct {
	// Name is the full name of the package in the repository
	Name string

	// Language is the programming language to generate code for
	Language Language

	// Source is the repository for protos
	Source source.Source

	// Ref defines a point in the repository timeline
	Ref source.Ref
}
