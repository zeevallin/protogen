package config

import (
	"fmt"

	"github.com/zeeraw/protogen/source"
)

const (
	fmtErrPrepare = "could not prepare package: %v"
	fmtErrClean   = "could not clean package: %v"
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

// CleanFunc is a type that represents a cleanup fucntion
type CleanFunc func() error

// Prepare the correct source by prepare out the branch
func (p *Package) Prepare() (CleanFunc, error) {
	clean := func() error {
		return nil
	}
	repo, err := p.Source.InitRepo()
	if err != nil {
		return clean, fmt.Errorf(fmtErrPrepare, err)
	}
	clean = func() error {
		if err := repo.Clean(); err != nil {
			return fmt.Errorf(fmtErrClean, err)
		}
		return nil
	}
	if err := repo.Checkout(p.Ref); err != nil {
		return clean, fmt.Errorf(fmtErrPrepare, err)
	}
	return clean, nil
}

// Path returns the absolute path to the package
func (p *Package) Path() string {
	return p.Source.PathTo(p.Name)
}

// Root returns the absolute path to the package import root
func (p *Package) Root() string {
	return p.Source.Root()
}
