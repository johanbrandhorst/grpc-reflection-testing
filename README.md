# grpc-reflection-testing
Playing around with gRPC reflection

## Usage

Install the grpcurl client

```bash
$ go install ./vendor/github.com/fullstorydev/grpcurl/cmd/grpcurl/
```

Start the server

```bash
$ go run main.go
```

List the services (this errors at the moment)

```bash
$ grpcurl -plaintext localhost:10000 list example.UserService
Failed to list methods for service "example.UserService": Symbol not found: example.UserService
caused by: File not found: github.com/mwitkow/go-proto-validators/validator.proto
```
