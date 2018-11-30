package protoc

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/zeeraw/protogen/config"
)

const (
	// Binary is the name of the protoc binary
	Binary = "protoc"

	versionFlag = "--version"
)

// Run starts a new protoc command line instance and runs a command against it
func Run(pkg *config.Package, files ...string) error {
	return NewProtoc().Run(pkg, files...)
}

// NewProtoc returns a new protoc instance
func NewProtoc() *Protoc {
	return &Protoc{
		WorkingDirectory: "",
		Binary:           Binary,
		Verbose:          false,
	}
}

// Protoc is the protoc biary
type Protoc struct {
	WorkingDirectory string
	Binary           string
	Verbose          bool
}

// Run will run the protoc command with
func (p *Protoc) Run(pkg *config.Package, files ...string) error {
	switch pkg.Language {
	case config.Go:
		return p.runGo(pkg, files...)
	}
	return ErrUnknownLanguage{pkg.Language}
}

// Exec will run commands against the protoc binary
func (p *Protoc) Exec(args ...string) error {
	command := exec.Command(p.Binary, args...)
	command.Dir = p.WorkingDirectory

	stderr, err := command.StderrPipe()
	if err != nil {
		return ErrFailedProtocExec{err.Error()}
	}
	if err := command.Start(); err != nil {
		return ErrFailedProtocExec{err.Error()}
	}

	out, err := ioutil.ReadAll(stderr)
	if err != nil {
		return ErrFailedProtocExec{err.Error()}
	}

	if err := command.Wait(); err != nil {
		return ErrFailedProtocExec{fmt.Sprintf("could not execute: %s: %v", out, err)}
	}
	return nil
}

// Test will check if protoc is installed on your system and give you an error if it isn't
func (p *Protoc) Test() (string, error) {
	command := exec.Command(p.Binary, versionFlag)
	out, err := command.Output()
	if err != nil {
		return "", ErrProtocMissing{err}
	}
	return string(out), nil
}

func (p *Protoc) log(msg string) {
	fmt.Printf("protogen: %s\n", msg)
}
