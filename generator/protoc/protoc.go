package protoc

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/zeeraw/protogen/config"
)

const (
	// Binary is the name of the protoc binary
	Binary = "protoc"

	versionFlag = "--version"
)

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
	log.Println("protoc selecting language")
	switch pkg.Language {
	case config.Go:
		return p.RunGo(pkg, files...)
	case config.Swift:
		return p.RunSwift(pkg, files...)
	default:
		return p.RunGeneric(pkg, files...)
	}
}

// Exec will run commands against the protoc binary
func (p *Protoc) Exec(args ...string) error {
	log.Printf("protoc executing: %s\n", strings.Join(args, " "))
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

// Check will look what protoc version is installed on your system
// Returns the version or error if it isn't installed
func (p *Protoc) Check() (string, error) {
	command := exec.Command(p.Binary, versionFlag)
	out, err := command.Output()
	if err != nil {
		return "", ErrProtocMissing{err}
	}
	return deriveVersion(string(out)), nil
}

// CheckExtension will look and see if an extension binary is installed
func (p *Protoc) CheckExtension(lang config.Language) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	binary := fmt.Sprintf("protoc-gen-%s", lang)
	out, err := exec.CommandContext(ctx, binary, versionFlag).Output()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError: // ExitError should trigger when the context timeouts
			return "", nil
		default:
			return "", ErrExtensionMissing{lang, err}
		}
	}
	return deriveVersion(string(out)), nil
}

func deriveVersion(s string) string {
	trimmed := strings.TrimSpace(string(s))
	tuple := strings.Split(trimmed, " ")
	return tuple[1]
}
