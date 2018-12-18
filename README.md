[![GoDoc](https://godoc.org/github.com/zeeraw/protogen?status.svg)](https://godoc.org/github.com/zeeraw/protogen)
[![Build Status](https://travis-ci.org/zeeraw/protogen.svg?branch=master)](https://travis-ci.org/zeeraw/protogen)

# protogen
Command line tool for centralised structuring of Google's protocol buffers.

## Motive
Using Google's protocol buffers bring great potential for substantial performance gains and type safety over the wire. But it also comes with the constraint that you need to share proto files among all your service providers and consumers.

Protogen works under the assumption, and with the approach that proto files should be centralised in one repository, divided into packages, version tagged and made available through Git. What the protogen tool provides is a workflow for organising and versioning your protobuf contracts by abstracting [**protoc**](https://github.com/protocolbuffers/protobuf) and [**git**](https://git-scm.com/), acting like a package manager for protoc generated code.

We want you to focus on collaborating with your colleuages, design powerful APIs and build great projects.

## Editor Support
- [Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=zeeraw.protogen) _(official)_

## Contributing
First of all, thanks for thinking about contributing to the protogen project. I think with a few people helping me make this tool a reality, it will save a lot of people a lof of trouble when managing their .proto files.

To get started you first need to pull the project into your `GOPATH` and change your directory to the repository.

```bash
$ git clone git@github.com:zeeraw/protogen.git $GOPATH/src/github.com/zeeraw/protogen
$ cd $GOPATH/src/github.com/zeeraw/protogen
```

### Installing dependencies
You need to have the [**protoc**](https://github.com/protocolbuffers/protobuf) tool installed on your computer, without it protogen will not work. If you got the protogen binary in your path, you can always run `protogen check` to see if all the dependencies are available.

After that you can run the install command and wait for it to complete. After that you should be all set to start adding your cool features or clear documentation to protogen.

```bash
# Installs dependencies and the protogen binary to your $GOPATH
$ make install
```

### Running tests
In the project we have two types of tests: regular tests and integration tests. We've made it fairly simple by providing three make targets.

```bash
# Runs only regular tests, should work without network access
$ make test
```

```bash
# Runs only integration tests, requires network access to pass
$ make integration
```

```bash
# Runs all tests, requires network access to pass
$ make all
```