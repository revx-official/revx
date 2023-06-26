# revx

---

Welcome to *revx*, a simple reverse proxy written in *Go*.

Features:

- proxy pass requests (http)
- round robin load balancing
- server health check routines
- simple configuration using yaml

Container Platforms:

- linux/amd64
- linux/arm/v5
- linux/arm/v7
- linux/arm64/v8

---

## Quickstart

For a quick start, use the official *revx* docker container:

```sh
$ docker pull ghcr.io/revx-official/revx:latest
$ docker run -p 9999:80 ghcr.io/revx-official/revx:latest
```

Alternatively, deploy *revx* instantly with docker compose:

```yaml
version: '3.8'

services:
  revx:
    image: ghcr.io/revx-official/revx:latest
    container_name: revx
    volumes:
      - ./config/config.yaml:/etc/revx/config.yaml
    ports:
      - "9999:80"
```

## Configuration

Example *revx* configuration:

```yaml
port: 80

servers:
  - name: server-1
    context: /one
    upstreams:
      - http://127.0.0.1:9991'
      - http://127.0.0.1:9992'
    allowed-methods:
      - GET
      - POST
      - PUT
      - DELETE
    health-check:
      endpoint: /
      interval: 10000
      fails: 5
  - name: server-2
    context: /two
    upstreams:
      - 'http://127.0.0.1:9993'
    allowed-methods:
      - GET
      - POST
      - PATCH
      - DELETE
    health-check:
      Endpoint: /health
      Interval: 5000
      Fails: 3
```

## Project Setup

To get *revx* up and running, follow the instructions below.

### Platforms

Officially supported development platforms are:

- Windows
- Linux
- Mac

### Go

The *revx* project is written in *Go*, hence it is required to install *Go*. For the latest version of *Go*, check: https://go.dev/doc/install

### Docker

To use *revx* with *docker*, you have to install *docker*, obviously.
For instructions, check out the official installation documentation: https://docs.docker.com/engine/install

## Getting Started

### Build

It is recommended to use the offical *revxbuildtool* to build *revx*. Install the tool by using the following command:

```sh
$ go install github.com/revx-official/revxbuildtool@latest
```

Build *revx* for local testing:

```sh
$ revxbuildtool --local
```

Build *revx* release version:

```sh
$ revxbuildtool --release
```

### Debugging

To debug *revx* use the officially provided `launch.json` for *Visual Studio Code*.

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/daemon.go",
            "buildFlags": "-tags local -ldflags '-X \"github.com/revx/pkg/revx.RevxApp=revx\" -X \"github.com/revx/pkg/revx.RevxModel=revx community\" -X \"github.com/revx/pkg/revx.RevxVersion=local\" -X \"github.com/revx/pkg/revx.RevxCommit=?\"'",
        }
    ],
}
```

Build the *revx* docker container:

```sh
$ docker build -t revx:local .
```

Using *revx* in docker compose:

```yaml
version: '3.8'

services:
  revx:
    image: revx:local
    container_name: revx
    volumes:
      - ./config/config.yaml:/etc/revx/config.yaml
    ports:
      - "9999:80"
```

## Version File

The `version.yaml` file contains metadata about the application, such as the current *revx* version. This metadata is compiled into the application by using *Golang's* compiler option `ldflags`.

```yaml
app: revx
model: <model>
version: <version>
```

## Documentation

For more information about *revx*, see the official documentation: [docs](./docs/index.md).