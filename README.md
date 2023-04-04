# gueue

A simple event queue powered by Go concurrency patterns and gRPC.

## Usage

### Installation

Server:

```shell
go install github.com/gaarutyunov/gueue/cmd/gueue
```

Client package:

```shell
go get github.com/gaarutyunov/gueue/pkg/client
```

### Configuration

An example configuration file for `gueue` server:

```yaml
server:
  host: localhost
  port: 8001

topics:
  - name: test.topic.1
    buffer: 1000
  - name: test.topic.2
    buffer: 1000
```

Start server with:
```shell
gueue server -c gueue.yaml
```

### Examples

For an example of client package usage see `examples` folder.

Start with a simple example in `examples/simple/consumer/main.go` and `examples/simple/producer/main.go`.

### Development

1. Build server
```shell
make build
```

2. Lint
```shell
make setup lint
```

3. Format
```shell
make fmt
```

4. Build `simple` example
```shell
make example name=simple
```
