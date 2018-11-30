// +build integration

package protoc

import (
	"os"
	"strings"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/test"

	"github.com/zeeraw/protogen/config/go"
)

func Test_Protoc_buildGo(t *testing.T) {
	wd, _ := os.Getwd()
	p := NewProtoc()
	p.WorkingDirectory = wd

	cfg := &config.Package{
		Output:   "./tmp",
		Language: config.Go,
		LanguageConfig: &golang.Config{
			Paths: golang.Import,
			Plugins: []golang.Plugin{
				golang.GRPC,
				golang.Plugin("test1"),
				golang.Plugin("test2"),
			},
		},
	}

	expected := []string{
		"--go_out=plugins=grpc+test1+test2,paths=import:./tmp",
		"hello/world.proto",
	}

	command, err := p.buildGo(cfg, "hello/world.proto")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, strings.Join(expected, " "), strings.Join(command, " "))
}

func Test_Protoc_runGo(t *testing.T) {
	wd, _ := os.Getwd()
	p := NewProtoc()
	p.WorkingDirectory = wd

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

	err := p.runGo(cfg, "fixtures/fixtures.proto")
	test.AssertEqual(t, nil, err)

}
