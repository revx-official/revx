# Proxy Passing

## Introduction

Proxy passing or proxy forwarding describes the action of passing a client request to one of the registered upstreams and passing the upstream response back to the client. The client never knows or cares that its request has been passed through a reverse proxy to the real server.

## Feature Support

At the moment *revx* supports forwarding any kind of HTTP request. Proxy forwarding with SSL/TLS is **not** supported currently.