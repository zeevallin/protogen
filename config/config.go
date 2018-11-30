package config

import (
	"fmt"

	"github.com/zeeraw/protogen/source"
)

const (
	errFmt = "could not prepare package: %v"
)

// Config defines the configuration for code generation
type Config struct {
	Packages []*Package
	General  *General
}

// General contains general config for generation
type General struct {
	Verbose bool
}

// Package defines a package to generate code for
type Package struct {
	// Name is the full name of the package in the repository
	Name string

	// Output is the output path on disk
	Output string

	// Language is the programming language to generate code for
	Language Language

	// LanguageConfig is the language specific configuration for generating code
	LanguageConfig interface{}

	// Source is the repository for protos
	Source source.Source

	// Ref defines a point in the repository timeline
	Ref source.Ref
}

// Prepare the correct source by prepare out the branch
func (p *Package) Prepare() error {
	err := p.Source.Init()
	if err != nil {
		return fmt.Errorf(errFmt, err)
	}
	hash, err := p.Source.HashForRef(p.Ref)
	if err != nil {
		return fmt.Errorf(errFmt, err)
	}
	err = p.Source.Checkout(hash)
	if err != nil {
		return fmt.Errorf(errFmt, err)
	}
	return p.Source.Checkout(hash)
}

// Path returns the absolute path to the package
func (p *Package) Path() string {
	return ""
}
