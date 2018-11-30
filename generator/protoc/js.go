package protoc

import (
	"fmt"

	"github.com/zeeraw/protogen/config"
)

const jsFlag = "--go_out"

func (p *Protoc) runJS(pkg *config.Package, files ...string) error {
	args, err := p.buildJS(pkg, files...)
	if err != nil {
		return err
	}
	return p.Exec(args...)
}

func (p *Protoc) buildJS(pkg *config.Package, files ...string) ([]string, error) {
	args := []string{}
	args = append(args, fmt.Sprintf("--js_out=library=%s,import_style=commonjs_strict,binary:%s", pkg.Name, pkg.Path()))
	args = append(args, files...)
	return args, nil
}
