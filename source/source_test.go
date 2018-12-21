package source_test

import (
	"os"
	"path"
	"testing"

	"github.com/zeeraw/protogen/source"
)

func TestMain(m *testing.M) {
	source.WorkDir = path.Join(os.TempDir(), "protogen-test")

	r := m.Run()
	defer os.Exit(r)

	err := os.RemoveAll(source.WorkDir)
	if err != nil {
		panic(err)
	}
}
