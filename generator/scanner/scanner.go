package scanner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	protoFileSuffix = ".proto"
)

// Scanner represents a .proto package scanner
type Scanner struct {
	GoPkgs map[string][]string
	Pkgs   map[string][]string

	root    string
	files   map[string]struct{}
	scanned map[string]struct{}
}

// New creates a new scanning instance
func New(root string) *Scanner {
	return &Scanner{
		files:   make(map[string]struct{}),
		scanned: make(map[string]struct{}),
		GoPkgs:  make(map[string][]string),
		Pkgs:    make(map[string][]string),
		root:    root,
	}
}

// Bundle is a bundle of
type Bundle struct {
	GoPackage string
	Packages  map[string]struct{}
}

// Scan will take a directory and recursively scan for proto files
func (s *Scanner) Scan(dict string) []Bundle {
	files := []string{}
	log.Printf("scanner scanning dictionary: %q\n", dict)
	filepath.Walk(dict, func(path string, f os.FileInfo, err error) error {
		if path != dict {
			if strings.HasSuffix(path, protoFileSuffix) {
				s.files[path] = struct{}{}
				files = append(files, path)
			}
		}
		return nil
	})

	err := s.extractFiles(files)
	if err != nil {
		panic(err)
	}

	bundles := []Bundle{}
	for goPkg, files := range s.GoPkgs {
		ss := Bundle{
			GoPackage: goPkg,
			Packages:  make(map[string]struct{}),
		}
		for _, file := range files {
			ss.Packages[path.Dir(file)] = struct{}{}
		}
		bundles = append(bundles, ss)
	}

	return bundles
}

type extraction struct {
	imports []string
	goPkg   string
	pkg     string
}

func (s *Scanner) extractFile(path string) (ex extraction, err error) {
	log.Printf("scanner extracting from file: %q\n", path)
	ex = extraction{
		imports: make([]string, 0),
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("scanner failed to open file: %q\n", s.shortPath(path))
		return ex, err
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		if strings.HasPrefix(scnr.Text(), "import") {
			txt := scnr.Text()
			txt = strings.TrimSuffix(txt, "\";")
			txt = strings.TrimPrefix(txt, "import \"")
			log.Printf("scanner extracting \"import\" in file %q: %s\n", s.shortPath(path), txt)
			ex.imports = append(ex.imports, s.joinPath(txt))
		}
		if strings.HasPrefix(scnr.Text(), "option go_package") {
			txt := scnr.Text()
			txt = strings.TrimSuffix(txt, "\";")
			txt = strings.TrimPrefix(txt, "option go_package = \"")
			log.Printf("scanner extracting \"go_package\" in file %q: %s\n", s.shortPath(path), txt)
			ex.goPkg = txt
		}
		if strings.HasPrefix(scnr.Text(), "package") {
			txt := scnr.Text()
			txt = strings.TrimSuffix(txt, ";")
			txt = strings.TrimPrefix(txt, "package ")
			log.Printf("scanner extracting \"package\" in file %q: %s\n", s.shortPath(path), txt)
			ex.pkg = txt
		}
	}

	if err := scnr.Err(); err != nil {
		return ex, err
	}

	s.scanned[path] = struct{}{}
	s.GoPkgs[ex.goPkg] = append(s.GoPkgs[ex.goPkg], path)
	s.Pkgs[ex.pkg] = append(s.Pkgs[ex.pkg], path)
	return ex, err
}

func (s *Scanner) extractFiles(paths []string) (err error) {
	for _, path := range paths {
		if _, ok := s.scanned[path]; !ok {
			ex, err := s.extractFile(path)
			if err != nil {
				continue
			}

			err = s.extractFiles(ex.imports)
			if err != nil {
				continue
			}
		}
	}
	return nil
}

func (s *Scanner) shortPath(p string) string {
	return strings.TrimPrefix(p, fmt.Sprintf("%s/", s.root))
}

func (s *Scanner) joinPath(p string) string {
	return path.Join(s.root, p)
}
