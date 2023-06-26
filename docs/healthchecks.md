# Health checks

## Introduction

What are health checks? A health check is a simple request to a server, to check whether this server still is available. Typically, health checks run periodically, to ensure, that a server is actually available.

## Health Checks In revx

revx health checks allow the configuration of 3 parameters. The `endpoint` refers to the actual path of the server, which is requested by the health check. As long as the request can be performed (a connection can be opened), the response of this endpoint doesn't matter.

The `interval` determines how many milliseconds to wait, until the next health check on this server is performed.

The `fails` parameter indicates after how many consecutive health check fails (the upstream is not reachable), the upstream is considered unhealthy. If an upstream is considered unhealthy, no more requests will be proxy forwarded to this upstream (except, there is only one upstream configured).

```yaml
health-check:
  endpoint: /health
  interval: 5000
  fails: 3
```