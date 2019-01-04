// +build integration

package protoc_test

import (
	"strings"
	"testing"

	"github.com/zeeraw/protogen/config"
	"github.com/zeeraw/protogen/config/swift"
	"github.com/zeeraw/protogen/test"
)

func Test_Protoc_BuildSwift(t *testing.T) {

	cases := []struct {
		name     string
		pkg      *config.Package
		files    []string
		expected []string
	}{
		{
			name: "valid public full path",
			pkg: &config.Package{
				Output:   "./tmp",
				Language: config.Swift,
				LanguageConfig: swift.Config{
					Options: []swift.Option{
						{
							Name:  "Visibility",
							Value: "Public",
						},
						{
							Name:  "FileNaming",
							Value: "FullPath",
						},
					},
				},
				Source: src,
			},
			files: []string{
				"hello/world.proto",
			},
			expected: []string{
				"--swift_opt=Visibility=Public",
				"--swift_opt=FileNaming=FullPath",
				"--swift_out=./tmp",
				"hello/world.proto",
			},
		},
		{
			name: "valid internal path to underscores",
			pkg: &config.Package{
				Output:   "./tmp",
				Language: config.Swift,
				LanguageConfig: swift.Config{
					Options: []swift.Option{
						{
							Name:  "Visibility",
							Value: "Internal",
						},
						{
							Name:  "FileNaming",
							Value: "PathToUnderscores",
						},
					},
				},
				Source: src,
			},
			files: []string{
				"foo/bar.proto",
				"baz/buz.proto",
			},
			expected: []string{
				"--swift_opt=Visibility=Internal",
				"--swift_opt=FileNaming=PathToUnderscores",
				"--swift_out=./tmp",
				"foo/bar.proto",
				"baz/buz.proto",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			command, err := p.BuildSwift(c.pkg, c.files...)
			test.AssertEqual(t, nil, err)
			test.AssertEqual(t, strings.Join(c.expected, " "), strings.Join(command[1:], " "))
		})
	}
}

func Test_Protoc_RunSwift(t *testing.T) {
	cfg := &config.Package{
		Output:   "./tmp",
		Language: config.Swift,
		LanguageConfig: swift.Config{
			Options: []swift.Option{
				{
					Name:  "Visibility",
					Value: "Public",
				},
				{
					Name:  "FileNaming",
					Value: "FullPath",
				},
			},
		},
		Source: src,
	}

	err := p.RunSwift(cfg, "fixtures/fixtures.proto")
	test.AssertEqual(t, nil, err)

}
