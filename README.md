# apple-dog

Small Go module scaffold for experiments and CLI tooling.

## Layout

Standard Go project layout:

```text
.
├── cmd/apple-dog/     # main package (binary entrypoint)
├── internal/greet/    # private library code (not importable by other modules)
├── go.mod
├── Makefile           # optional shortcuts; plain `go` commands work too
└── README.md
```

- **`cmd/`** — one directory per command/binary. Build with `go build -o bin/apple-dog ./cmd/apple-dog`.
- **`internal/`** — packages used only inside this module; the Go toolchain enforces that external modules cannot import them.

## Requirements

- Go 1.22+

## Build

```bash
go build ./...
```

Or via Makefile:

```bash
make build    # writes bin/apple-dog
```

## Run

```bash
go run ./cmd/apple-dog
go run ./cmd/apple-dog -name Alice
```

Or:

```bash
make run
./bin/apple-dog -name Alice
```

## Test

```bash
go test ./...
```

Or:

```bash
make test
```

## Module path

`github.com/liucong009/apple-dog`
