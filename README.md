[![GoDoc](https://godoc.org/github.com/zeeraw/protogen?status.svg)](https://godoc.org/github.com/zeeraw/protogen)
[![Build Status](https://travis-ci.org/zeeraw/protogen.svg?branch=master)](https://travis-ci.org/zeeraw/protogen)

# protogen
Command line tool and workflow for organising code generation of Google's protocol buffers across multiple projects.

## Purpose
Using Google's protocol buffers bring great potential for substantial performance gains and type safety over the wire. But it also comes with the constraint that you need to share proto files among all your service providers and consumers.

Protogen works under the assumption, and with the approach that proto files should be centralised in one repository, divided into packages, version tagged and made available through Git. What the protogen tool provides is a workflow for organising and versioning your protobuf contracts by abstracting [**protoc**](https://github.com/protocolbuffers/protobuf) and [**git**](https://git-scm.com/), acting like a package manager for protoc generated code.

Our motive is to help you focus on collaborating with your colleuages, design powerful APIs and build great projects.

## What is protogen?
The protogen project consists of three things, the [**command line tool**](https://github.com/zeeraw/protogen/wiki/Command-line-tool), the [**workflow**](https://github.com/zeeraw/protogen/wiki/Workflow) and the [**.protogen file**](https://github.com/zeeraw/protogen/wiki/Protogen-file) written with [**the protogen configuration language**](https://github.com/zeeraw/protogen/wiki/Configuration-language). These are all the things you need to manage your proto files.

### The command line tool
The command line tool is used to both manage your central protobuf repository and to automate code generation in service and consumer projects.

#### Installing
The easiest way of obtaining the protogen command line tool for any platform is by installing it through Go package installer. Run the command and immediately have access to the tool.

```bash
$ go get -u github.com/zeeraw/protogen
```

### The workflow
The protogen workflow is how you should manage your central protobuf repository. This includes how to structure your packages and how to tag your releases. Visit our [**example repository**](https://github.com/zeeraw/protogen-protos) where we showcase how to structure our proto files. You can also [**read more about the protogen workflow on the wiki**](https://github.com/zeeraw/protogen/wiki/Workflow).

### The .protogen file
[**The .protogen file**](https://github.com/zeeraw/protogen/wiki/Protogen-file) is a definition of what protobuf code should be generated inside a given project. To write the file you use [**the protogen configuration language**](https://github.com/zeeraw/protogen/wiki/Configuration-language).

```protogen
source github.com/zeeraw/protogen-protos
output ./vendor/protos
language go {
    plugin grpc
}

generate services/games v1.0.1
generate services/foobar v3.0.0
```

#### Editor support
We provide protogen file syntax highlighting and snippets for a few different editors.
- [**Visual Studio Code**](https://marketplace.visualstudio.com/items?itemName=zeeraw.protogen) _(official)_
