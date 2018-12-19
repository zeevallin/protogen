# Contributing
First of all, thanks for thinking about contributing to the protogen project. We think that with a few people helping us make this tool a reality, it will save many people a lof of time and effort when bringing protobuffers into their organisations.

To get started you first need to pull the project into your `GOPATH` and change your directory to the repository.

```bash
$ git clone git@github.com:zeeraw/protogen.git $GOPATH/src/github.com/zeeraw/protogen
$ cd $GOPATH/src/github.com/zeeraw/protogen
```

## The workboard
We currently have a [**workboard for minimum viable product**](https://github.com/zeeraw/protogen/projects/1) where we have all the bugs, optimisations and features we want before making a first release. Go have a look and see if there's something there you would be able to do.

## Installing dependencies
You need to have the [**protoc**](https://github.com/protocolbuffers/protobuf) tool installed on your computer, without it protogen will not work. If you got the protogen binary in your path, you can always run `protogen info` to see if all the dependencies are available.

After that you can run the install command and wait for it to complete. After that you should be all set to start adding your cool features or clear documentation to protogen.

```bash
# Installs dependencies and the protogen binary to your $GOPATH
$ make install
```

## Running tests
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
