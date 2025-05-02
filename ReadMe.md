# Project Description

This is a basic **echo-server** running on TCP port 9000, built with Go and the **gnet v2** library. The server listens for connections from clients, receives data and returns the same content (echo) asynchronously, taking advantage of multicore for optimal performance.

## Features

* Uses **gnet v2** (event-driven, high-performance networking framework)
* Supports **multicore** automatically, utilizing all CPUs
* Handles asynchronous, non-blocking connections (Async I/O)
* Explicit callbacks: OnBoot, OnOpen, OnTraffic, OnClose, OnTick, OnShutdown

## Requirements

* Go 1.18+
* Go module (go.mod)
* ​​gnet v2 (`github.com/panjf2000/gnet/v2`)

## Installation

1. Initialize the module:

```bash
go mod init gnet-server-test
```
2. Load dependencies:

```bash
go mod tidy
```
## How to run

```bash
# Compile
go run main.go

```

The default server listens on `0.0.0.0:9000`.

## Test with ncat

1. Open a terminal, connect to the server:

```bash
ncat 127.0.0.1 9000
```
2. Type a line and press Enter, for example:

```
Hello gnet!
```
3. You should get a response:

```
Server received: Hello gnet!
```

Or automatically:

```bash
echo "Automatic Test" | ncat 127.0.0.1 9000
```

## Customization

* Change port: edit `"tcp://0.0.0.0:9000"` in `gnet.Run`
* Add processing logic in `OnTraffic` to build real applications (chat server, game server...)

## License

MIT © vancong230502
