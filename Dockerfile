FROM golang:latest
RUN apt-get update
RUN apt-get install -y protobuf-compiler
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get github.com/golang/mock/gomock
RUN go get github.com/golang/mock/mockgen
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . . 
RUN go generate ./...
RUN go build -o main . 
ENTRYPOINT ["/app/main", "-s"]
