package protoc

import (
	"fmt"
	"log"
	"strings"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/go"
)

const goFlag = "--go_out"

// RunGo will run generation of a package with Go code as the output
func (p *Protoc) RunGo(pkg *config.Package, files ...string) error {
	log.Println("protoc running go")
	args, err := p.BuildGo(pkg, files...)
	if err != nil {
		return err
	}
	return p.Exec(args...)
}

// BuildGo will construct the protoc command for a package with Go code as the output
func (p *Protoc) BuildGo(pkg *config.Package, files ...string) ([]string, error) {
	log.Println("protoc building go")
	cfg, ok := pkg.LanguageConfig.(golang.Config)
	if !ok {
		return nil, ErrConfigType{pkg.LanguageConfig}
	}

	lang := []string{}
	lang = append(lang, fmt.Sprintf("%s=", goFlag))
	var extras = []string{}
	if len(cfg.Plugins) > 0 {
		plugins := []string{}
		plugins = append(plugins, "plugins=")
		plugins = append(plugins, strings.Join(cfg.Plugins, "+"))
		extras = append(extras, strings.Join(plugins, ""))
	}
	if cfg.Paths != "" {
		extras = append(extras, fmt.Sprintf("paths=%s", cfg.Paths))
	}
	if len(extras) > 0 {
		lang = append(lang, strings.Join(extras, ","))
		lang = append(lang, ":")
	}
	lang = append(lang, pkg.Output)

	args := []string{}
	args = append(args, fmt.Sprintf("-I%s", pkg.Root()))
	args = append(args, strings.Join(lang, ""))
	args = append(args, files...)
	return args, nil
}
