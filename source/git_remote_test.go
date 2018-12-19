// +build integration

package source_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/zeeraw/protogen/source"
	"github.com/zeeraw/protogen/test"
)

const (
	masterCommitHash           = "5deeaaf1bfd117031b24e55182acdae386c14941"
	servicesFoobarCommitHashV1 = "4d2b63a1aef7c1ba5f6ff220005d9f4d8ea94443"
)

func TestRemoteGitSource(t *testing.T) {
	src, err := source.NewRemoteGitSource(log.New(ioutil.Discard, "", 0), "github.com/zeeraw/protogen-protos")
	test.AssertEqual(t, nil, err)

	err = src.Init()
	test.AssertEqual(t, nil, err)

	hash, err := src.HashForRef(source.Ref{
		Name: "master",
		Type: source.Branch,
	})
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, masterCommitHash, hash)

	err = src.Checkout(hash)
	test.AssertEqual(t, nil, err)

	hash, err = src.HashForRef(source.Ref{
		Name: "services/foobar/v1.0.0",
		Type: source.Version,
	})
	test.AssertEqual(t, nil, err)
	test.AssertEqual(t, servicesFoobarCommitHashV1, hash)

	err = src.Checkout(hash)
	test.AssertEqual(t, nil, err)
}
