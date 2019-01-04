package config_test

import (
	"fmt"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/source"
)

func TestConfig(t *testing.T) {
	t.Run("can create", func(tt *testing.T) {
		cfg := config.Config{
			General: &config.General{
				Verbose: false,
			},
			Packages: []*config.Package{
				{
					Name:     "master",
					Language: "go",
					Source:   source.NewMockGitSource(""),
					Ref:      source.Ref{},
				},
			},
		}
		fmt.Println(cfg)
	})
}
