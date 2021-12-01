# EdgeStats Server

The [EdgeStats](https://www.edgestats.io) server. May be used with private [edgestats-client](https://github.com/edgestats/edgestats-client) and [edgestats-webui](https://github.com/edgestats/edgestats-webui) setup.
>
> Check your Theta Edge Node uptime stats without needing to view the edge node client GUI/CLI.

## Setup

### Clone repository
Execute the following commands to clone this repository:

```shell
git clone https://github.com/edgestats/edgestats-server
cd edgestats-server
```

### Install dependencies
Execute the following command to install the dependencies:

```shell
go mod tidy
```

### Build server from source
```shell
GOOS=<OS> GOARCH=<ARCH> go build -ldflags "-X 'main.srvPort=<port>' -X 'github.com/edgestats/edgestats-server/handlers.apiKey=<your-api-key>'" -o ./build/edgestats-server-<OS>-<ARCH> ./cmd/main.go
# example: GOOS=linux GOARCH=amd64 go build -ldflags "-X 'main.srvPort=8000' -X 'github.com/edgestats/edgestats-server/handlers.apiKey=thetaverse'" -o ./build/edgestats-server-linux-amd64 ./cmd/main.go
```

### Start server
```shell
./build/edgestats-server-<OS>-<ARCH>
# example: ./build/edgestats-server-linux-amd64
```

### Setup EdgeStats client (see Advanced Setup)
Instructions for setting up an EdgeStats client available [here](https://github.com/edgestats/edgestats-client).

### Setup EdgeStats webui
Instructions for setting up an EdgeStats webui available [here](https://github.com/edgestats/edgestats-webui).

## LICENSE
Copyright (c) EdgeStats Authors