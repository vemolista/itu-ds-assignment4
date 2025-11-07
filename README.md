# ITU Distributed Systems - Assignment 4

## Running the application

Generate proto code
```sh
protoc --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/proto.proto
```

Compile
```sh
go build
```

Run three nodes (in separate terminals)
```sh
# Terminal 1
./itu-ds-assignment4 -i 0

# Terminal 2
./itu-ds-assignment4 -i 1

# Terminal 3
./itu-ds-assignment4 -i 2
```

Observe logs in `./app.log`