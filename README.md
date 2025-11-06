# ITU Distributed Systems - Assignment 4

## Starting the application

Generate proto code
```sh
protoc --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/proto.proto
```