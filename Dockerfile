FROM golang:latest
RUN mkdir /app 
ENV PATH=$GOPATH/bin:$PATH
RUN echo $PATH
WORKDIR /app 
RUN apt-get update
RUN apt-get install -y protobuf-compiler
COPY go.mod .
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go mod download
COPY . . 
RUN go generate ./...
RUN go test ./...
RUN go build -o main . 
ENTRYPOINT ["/app/main", "-s"]
