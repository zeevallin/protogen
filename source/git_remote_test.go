// +build integration

package source_test

import (
	"testing"

	"github.com/zeeraw/protogen/source"
	"github.com/zeeraw/protogen/test"
)

const (
	masterCommitHash           = "c3ca99b15b140f7ee346d4dad585501215dd6d48"
	servicesFoobarCommitHashV1 = "4d2b63a1aef7c1ba5f6ff220005d9f4d8ea94443"
)

func TestRemoteGitSource(t *testing.T) {
	src, err := source.NewRemoteGitSource("github.com/zeeraw/protogen-protos")
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
