# Todo Backend

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

There are two probes available:

- `/livez`
- `/readyz`
