package config_test

import (
	"fmt"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/source"
)

func ConfigTest(t *testing.T) {
	t.Run("can create", func(tt *testing.T) {
		cfg := config.Config{
			Packages: []*config.Package{
				{
					Name:     "master",
					Language: "go",
					Source:   &source.RemoteGitSource{},
					Ref:      source.Ref{},
				},
			},
		}
		fmt.Println(cfg)
	})
}
