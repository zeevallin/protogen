package protoc

import (
	"fmt"
	"strings"

	"github.com/zeeraw/protogen/config"
)

// RunGeneric will run generation of a package in a generic way
func (p *Protoc) RunGeneric(pkg *config.Package, files ...string) error {
	p.logger.Printf("protoc running generic for %q\n", pkg.Language)
	args, err := p.BuildGeneric(pkg, files...)
	if err != nil {
		return err
	}
	return p.Exec(args...)
}

// BuildGeneric will construct the protoc command for a package in a generic way
func (p *Protoc) BuildGeneric(pkg *config.Package, files ...string) ([]string, error) {
	p.logger.Printf("protoc building generic for %q\n", pkg.Language)
	flag := fmt.Sprintf("--%s_out", pkg.Language)
	lang := []string{}
	lang = append(lang, fmt.Sprintf("%s=", flag))
	lang = append(lang, pkg.Output)

	args := []string{}
	args = append(args, fmt.Sprintf("-I%s", pkg.Root()))
	args = append(args, strings.Join(lang, ""))
	args = append(args, files...)
	return args, nil
}
