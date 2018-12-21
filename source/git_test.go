// +build integration

package source_test

import (
	"path"
	"runtime"
	"testing"

	"github.com/zeeraw/protogen/source"
	"github.com/zeeraw/protogen/test"
)

const (
	masterCommitHash           = "5deeaaf1bfd117031b24e55182acdae386c14941"
	servicesFoobarCommitHashV1 = "4d2b63a1aef7c1ba5f6ff220005d9f4d8ea94443"
)

func TestRemoteGitSource(t *testing.T) {
	src, err := source.NewRemoteGitSource("github.com/zeeraw/protogen-protos")
	test.AssertEqual(t, nil, err)

	_, err = src.InitRepo()
	test.AssertEqual(t, nil, err)
}

func TestLocalGitSource(t *testing.T) {
	src, err := source.NewLocalGitSource(fixtures())
	test.AssertEqual(t, nil, err)

	_, err = src.InitRepo()
	test.AssertEqual(t, nil, err)
}

func fixtures() string {
	_, file, _, _ := runtime.Caller(0)
	dir := path.Dir(file)
	return path.Join(dir, "fixtures")
}
