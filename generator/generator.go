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
	logger  *log.Logger
}

// Generate generates based on a protogen project configuration
func Generate(logger *log.Logger, cfg *config.Config) error {
	logger.Printf("generating config: %d packages", len(cfg.Packages))
	for _, pkg := range cfg.Packages {
		logger.Printf("generating package: %s", pkg.Name)
		if err := GeneratePackage(logger, pkg, cfg.General); err != nil {
			return err
		}
	}
	return nil
}

// GeneratePackage generates a specific package and its dependencies
func GeneratePackage(logger *log.Logger, pkg *config.Package, general *config.General) error {
	return NewGenerator(logger, pkg, general).Run()
}

// NewGenerator constructs a code generator with config
func NewGenerator(logger *log.Logger, pkg *config.Package, general *config.General) *Generator {
	return &Generator{
		pkg:     pkg,
		general: general,
		logger:  logger,
	}
}

// Run will generate the
func (g *Generator) Run() error {
	g.logger.Printf("generator running\n")
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	p := protoc.NewProtoc(g.logger)
	p.WorkingDirectory = wd

	g.logger.Printf("generator preparing package\n")
	clean, err := g.pkg.Prepare()
	defer clean()
	if err != nil {
		return err
	}

	g.logger.Printf("generator scanning\n")
	scanner := scanner.New(g.logger, g.pkg.Root())
	scanner.Scan(g.pkg.Path())

	switch g.pkg.Language {
	case config.Go:
		g.logger.Printf("generator scanned %d go packages\n", len(scanner.GoPkgs))
		for goPkg, files := range scanner.GoPkgs {
			for _, f := range files {
				name := strings.TrimPrefix(f, g.pkg.Source.PathTo(g.pkg.Name))
				g.logger.Printf("generating %s for go package \"%s\" into %s", name, goPkg, g.pkg.Output)
			}
			if err := p.Run(g.pkg, files...); err != nil {
				return err
			}
		}
	default:
	}

	return nil
}
