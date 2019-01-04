package protoc

import (
	"fmt"
	"log"
	"strings"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/swift"
)

const (
	swiftOutFlag    = "--swift_out"
	swiftOptFlag    = "--swift_opt"
	swiftVisibility = "Visibility"
	swiftFileNaming = "FileNaming"
	swiftMappings   = "ProtoPathModuleMappings"
)

// RunSwift will run generation of a package with Swift code as the output
func (p *Protoc) RunSwift(pkg *config.Package, files ...string) error {
	log.Println("protoc running swift")
	args, err := p.BuildSwift(pkg, files...)
	if err != nil {
		return err
	}
	return p.Exec(args...)
}

// BuildSwift will construct the protoc command for a package with Swift code as the output
func (p *Protoc) BuildSwift(pkg *config.Package, files ...string) ([]string, error) {
	log.Println("protoc building swift")
	cfg, ok := pkg.LanguageConfig.(swift.Config)
	if !ok {
		return nil, ErrConfigType{pkg.LanguageConfig}
	}
	lang := []string{}
	lang = append(lang, fmt.Sprintf("%s=", swiftOutFlag))
	lang = append(lang, pkg.Output)

	options := []string{}
	for _, option := range cfg.Options {
		options = append(options, buildSwiftOption(option.Name, option.Value))
	}

	args := []string{}
	args = append(args, fmt.Sprintf("-I%s", pkg.Root()))
	for _, option := range options {
		args = append(args, option)
	}
	args = append(args, strings.Join(lang, ""))
	args = append(args, files...)
	return args, nil
}

func buildSwiftOption(option, value string) string {
	return fmt.Sprintf("%s=%s=%s", swiftOptFlag, option, value)
}
