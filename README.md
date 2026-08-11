 
# nofreedom

A simple command-line unit converter written in Go. It supports common metric
and imperial units for length and mass.

## Requirements

- Go 1.26.5 or newer

## Run

Run the CLI directly from the project directory:

```sh
go run ./cmd/nofreedom --help
```

Convert a value by passing the value, source unit, and target unit:

```sh
go run ./cmd/nofreedom convert 10 m ft
go run ./cmd/nofreedom convert 5 kg lb
```

List all supported units:

```sh
go run ./cmd/nofreedom list
```

## Build

```sh
go build -o nofreedom ./cmd/nofreedom
./nofreedom convert 10 m ft
```

## Supported units

- Length: `m`, `km`, `cm`, `mm`, `ft`, `mi`, `in`, `yd`
- Mass: `kg`, `g`, `mg`, `lb`, `oz`

Conversions must use units from the same category.

## Tests

```sh
go test ./...
```
