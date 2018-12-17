// +build integration

package protoc_test

import (
	"strings"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/test"

	"github.com/zeeraw/protogen/config/go"
)

func Test_Protoc_BuildGo(t *testing.T) {
	cfg := &config.Package{
		Output:   "./tmp",
		Language: config.Go,
		LanguageConfig: golang.Config{
			Paths: golang.Import,
			Plugins: []golang.Plugin{
				golang.GRPC,
				golang.Plugin("test1"),
				golang.Plugin("test2"),
			},
		},
		Source: src,
	}

	expected := []string{
		"--go_out=plugins=grpc+test1+test2,paths=import:./tmp",
		"hello/world.proto",
	}

	command, err := p.BuildGo(cfg, "hello/world.proto")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, strings.Join(expected, " "), strings.Join(command[1:], " "))
}

func Test_Protoc_RunGo(t *testing.T) {
	cfg := &config.Package{
		Output:   "./tmp",
		Language: config.Go,
		LanguageConfig: golang.Config{
			Paths: golang.Import,
			Plugins: []golang.Plugin{
				golang.GRPC,
			},
		},
		Source: src,
	}

	err := p.RunGo(cfg, "fixtures/fixtures.proto")
	test.AssertEqual(t, nil, err)

}
