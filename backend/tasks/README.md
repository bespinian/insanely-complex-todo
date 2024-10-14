# ICT - Tasks API

This is the tasks backend for our Todo app. It's a simple [Fiber](https://docs.gofiber.io/) API which has one resource: Tasks.

If nothing else is configured, it just stores the tasks in-memory. This means, the tasks are gone, when the server is restarted.

You can set the environment variable `MONGODB_URI` to define the URI for a MongoDB connection. If set, tasks are persistently stored in MongoDB and not in-memory.

## Build

```sh
go build main.go
```

## Development

```sh
go install github.com/mitranim/gow@latest
gow -c -e=go,mod,sum,env,toml run main.go
```

## Probes

There are two probes available: See [Fiber's health check](https://docs.gofiber.io/api/middleware/healthcheck) documentation for more details.

- `/livez`
- `/readyz`
