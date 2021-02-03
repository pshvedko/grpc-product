# grpc-product-demo
MongoDB based gRPC service that provides Fetch and List methods

## Installation
Get sources from repository  
```
git clone https://github.com/pshvedko/grpc-product.git
cd grpc-product
```

### Server
Build server in docker and start it with mongodb and nginx in composer
```
docker-compose up
```

### Client
Build client locally
```
apt-get update
apt-get install protobuf-compiler
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
export PATH=$GOPATH/bin:$PATH
go generate ./...
go build
```

### Usage
#### Help
``` 
./grpc-product -h
```
```
Usage:
  grpc-product [flags]
  grpc-product [command]

Available Commands:
  fetch       Load external CSV file
  help        Help about any command
  list        Browses the contents of storage

Flags:
  -s, --              run in service mode
  -a, --addr ip       address to bind (default 0.0.0.0)
  -h, --help          help for grpc-product
  -n, --node uint32   node id used with -s
  -p, --port int      port to listen (default 9000)

Use "grpc-product [command] --help" for more information about a command.

```
#### Fetch
Getting CSV file and saving to mongodb
```
./grpc-product fetch http://host/file.csv
```
#### List
Getting a list of products of 5 lines with an offset of 10 lines when sorting 
by changes and reverse sorting by price
```
./grpc-product list -o 10 -l 5 -s changes -s -price 
```
