# Configuration

The *revx* configuration is structured as follows. The `port` parameter specifies the actual port, on which *revx* itself runs on. The `servers` list specifies all servers to which *revx* should proxy forward to. A server always has a `context`. This parameter specifes for which URL path a request is forwarded to a server. It is important to note, that the context path is always included when proxy forwarding. So, if the context of a server is `/one` and the upstream is `http://127.0.0.1:9991`, a request is automatically forwarded to `http://127.0.0.1:9991/one`.

The `allowed-methods` describe all HTTP methods which are allowed and forwarded to the server.

The `health-check` properties describe how the internal health check routine for this server behaves. For more details on health checks, see [here](./healthchecks.md).


## Example

Here's an example configuration to get started:

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
