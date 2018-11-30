// +build integration

package protoc_test

import (
	"os"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/go"
	"github.com/zeeraw/protogen/generator/protoc"
	"github.com/zeeraw/protogen/test"
)

var (
	p *protoc.Protoc
)

func TestMain(m *testing.M) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p = protoc.NewProtoc()
	p.WorkingDirectory = wd

	os.Exit(m.Run())
}

func TestProtoc_Test(t *testing.T) {
	lib, err := p.Test()
	test.AssertEqual(t, nil, err)
	if len(lib) < 1 {
		t.Fatalf("lib should be more than 1 character: %s", lib)
	}
}

func TestProtoc_Run(t *testing.T) {
	cfg := &config.Package{
		Output:   "./tmp",
		Language: config.Go,
		LanguageConfig: &golang.Config{
			Paths: golang.Import,
			Plugins: []golang.Plugin{
				golang.GRPC,
			},
		},
	}

	err := p.Run(cfg, "fixtures/fixtures.proto")
	test.AssertEqual(t, nil, err)
}
