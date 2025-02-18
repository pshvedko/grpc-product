FROM golang:latest
RUN apt-get update
RUN apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/golang/mock/mockgen@latest
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . . 
RUN go generate ./...
RUN go build -o main . 
ENTRYPOINT ["/app/main", "-s"]
