// +build integration

package protoc_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/go"
	"github.com/zeeraw/protogen/generator/protoc"
	"github.com/zeeraw/protogen/source"
	"github.com/zeeraw/protogen/test"
)

var (
	logger *log.Logger
	p      *protoc.Protoc
	src    source.Source
)

func TestMain(m *testing.M) {
	logger = log.New(ioutil.Discard, "test", 0)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p = protoc.NewProtoc(logger)
	p.WorkingDirectory = wd

	src, err = source.NewMockSource(logger, "./")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestProtoc_Check(t *testing.T) {
	ver, err := p.Check()
	test.AssertEqual(t, nil, err)
	if len(ver) != 5 {
		t.Fatalf("version should be 5 characters, was %d: %q", len(ver), ver)
	}
}
func TestProtoc_CheckExtension(t *testing.T) {
	t.Run("when the extension is installed", func(t *testing.T) {
		err := p.CheckExtension("go")
		test.AssertEqual(t, nil, err)
	})
	t.Run("when the extension is not installed", func(t *testing.T) {
		s := "mumbojumbo"
		err := p.CheckExtension(s)
		if err != nil {
			switch err.(type) {
			case protoc.ErrExtensionMissing:
			default:
				t.Errorf("wrong type of error: %v", err)
			}
			return
		}
		t.Errorf("check extension with succeeded: %q", s)
	})
}

func TestProtoc_Run(t *testing.T) {
	cfg := &config.Package{
		Output:   "./tmp",
		Source:   src,
		Language: config.Go,
		LanguageConfig: golang.Config{
			Paths: golang.Import,
			Plugins: []golang.Plugin{
				golang.GRPC,
			},
		},
	}

	err := p.Run(cfg, "fixtures/fixtures.proto")
	test.AssertEqual(t, nil, err)
}
