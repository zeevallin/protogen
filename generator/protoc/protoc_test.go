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
	lib, err := p.Check()
	test.AssertEqual(t, nil, err)
	if len(lib[0]) < 1 {
		t.Fatalf("lib name be more than 1 character: %s", lib)
	}
	if len(lib[1]) < 1 {
		t.Fatalf("lib version be more than 1 character: %s", lib)
	}
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
