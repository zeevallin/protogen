package generator

import (
	"log"
	"os"
	"strings"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/generator/protoc"
	"github.com/zeeraw/protogen/generator/scanner"
)

// Generator defines the code generation session for one package and its dependent packages
type Generator struct {
	pkg     *config.Package
	general *config.General
}

// Generate generates based on a protogen project configuration
func Generate(cfg *config.Config) error {
	log.Printf("generating config: %d packages", len(cfg.Packages))
	for _, pkg := range cfg.Packages {
		log.Printf("generating package: %s", pkg.Name)
		if err := GeneratePackage(pkg, cfg.General); err != nil {
			return err
		}
	}
	return nil
}

// GeneratePackage generates a specific package and its dependencies
func GeneratePackage(pkg *config.Package, general *config.General) error {
	return NewGenerator(pkg, general).Run()
}

// NewGenerator constructs a code generator with config
func NewGenerator(pkg *config.Package, general *config.General) *Generator {
	return &Generator{
		pkg:     pkg,
		general: general,
	}
}

// Run will generate the
func (g *Generator) Run() error {
	log.Printf("generator running\n")
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	p := protoc.NewProtoc()
	p.WorkingDirectory = wd

	log.Printf("generator preparing package\n")
	clean, err := g.pkg.Prepare()
	defer clean()
	if err != nil {
		return err
	}

	log.Printf("generator scanning\n")
	scanner := scanner.New(g.pkg.Root())
	scanner.Scan(g.pkg.Path())

	switch g.pkg.Language {
	case config.Go:
		log.Printf("generator scanned %d go packages\n", len(scanner.GoPkgs))
		for goPkg, files := range scanner.GoPkgs {
			for _, f := range files {
				name := strings.TrimPrefix(f, g.pkg.Source.PathTo(g.pkg.Name))
				log.Printf("generating %s for go package \"%s\" into %s", name, goPkg, g.pkg.Output)
			}
			if err := p.Run(g.pkg, files...); err != nil {
				return err
			}
		}
	default:
		log.Printf("generator scanned %d packages\n", len(scanner.Pkgs))
		for pkg, files := range scanner.Pkgs {
			for _, f := range files {
				name := strings.TrimPrefix(f, g.pkg.Source.PathTo(g.pkg.Name))
				log.Printf("generating %s for package \"%s\" into %s", name, pkg, g.pkg.Output)
			}
			if err := p.Run(g.pkg, files...); err != nil {
				return err
			}
		}
	}

	return nil
}
