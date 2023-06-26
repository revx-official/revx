# Build

## Introduction

*revx* uses a custom build procedure to embed variables at compile time. To build *revx* the right way, use the official build tool: [revxbuildtool](https://github.com/revx-official/revxbuildtool)

## Installation

To install the build tool, use the `go install` command:

```sh
$ go install github.com/revx-official/revxbuildtool@latest
```

## Usage

To build *revx* for local testing, use:

```sh
$ revxbuildtool --local
```

To build a release version of *revx*, use one of the following commands:

```sh
# simply
$ revxbuildtool

# or
$ revxbuildtool --release
```
