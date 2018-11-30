package protoc

import (
	"fmt"
	"strings"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/go"
)

const goFlag = "--go_out"

func (p *Protoc) runGo(pkg *config.Package, files ...string) error {
	p.logger.Println("protoc running go")
	args, err := p.buildGo(pkg, files...)
	if err != nil {
		return err
	}
	return p.Exec(args...)
}

func (p *Protoc) buildGo(pkg *config.Package, files ...string) ([]string, error) {
	p.logger.Println("protoc building go")
	cfg, ok := pkg.LanguageConfig.(*golang.Config)
	if !ok {
		return nil, ErrConfigType
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
	if len(extras) > 1 {
		lang = append(lang, strings.Join(extras, ","))
		lang = append(lang, ":")
	}
	lang = append(lang, pkg.Output)

	args := []string{}
	args = append(args, strings.Join(lang, ""))
	args = append(args, files...)
	return args, nil
}
