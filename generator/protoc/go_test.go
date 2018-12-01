// +build integration

package protoc

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/zeeraw/protogen/source"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/test"

	"github.com/zeeraw/protogen/config/go"
)

func Test_Protoc_buildGo(t *testing.T) {
	wd, _ := os.Getwd()
	logger := log.New(ioutil.Discard, "test", 0)
	p := NewProtoc(logger)
	p.WorkingDirectory = wd
	s, err := source.NewRemoteGitSource(logger, "")
	if err != nil {
		t.Fatal(err)
	}
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
		Source: s,
	}

	expected := []string{
		"--go_out=plugins=grpc+test1+test2,paths=import:./tmp",
		"hello/world.proto",
	}

	command, err := p.buildGo(cfg, "hello/world.proto")
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, strings.Join(expected, " "), strings.Join(command[1:], " "))
}

func Test_Protoc_runGo(t *testing.T) {
	wd, _ := os.Getwd()
	logger := log.New(ioutil.Discard, "test", 0)
	p := NewProtoc(logger)
	p.WorkingDirectory = wd
	s, err := source.NewRemoteGitSource(logger, "")
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Package{
		Output:   "./tmp",
		Language: config.Go,
		LanguageConfig: golang.Config{
			Paths: golang.Import,
			Plugins: []golang.Plugin{
				golang.GRPC,
			},
		},
		Source: s,
	}

	err = p.runGo(cfg, "fixtures/fixtures.proto")
	test.AssertEqual(t, nil, err)

}
