package config

import (
	"github.com/zeeraw/protogen/source"
)

// Config defines the configuration for code generation
type Config struct {
	Projects []*Project
}

// Project defines a project to generate code for
type Project struct {
	Target   string
	Tag      string
	Language Language
	Source   source.Source
}
