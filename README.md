# FULLCYCLE: Comunicação entre sistemas - gRPC

## gRPC Project
This project demonstrates a simple implementation of gRPC services in Go.

## Features

## Installation
**1 - Install Protocol buffer compiler:**

$ apt install -y protobuf-compiler 

$ protoc --version  # Ensure compiler version is 3+

https://grpc.io/docs/protoc-installation/

**2 - Install protocol compiler plugins for GO**

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

**3 - Update path**

$ export PATH="$PATH:$(go env GOPATH)/bin"

**4 - Dependencies**

$ go mod tidy

**5 - Generate protobuf files**

protoc --go_out=. --go-grpc_out=. proto/course_category.proto

### Run



### Test

