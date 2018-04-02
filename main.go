package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	pbExample "github.com/johanbrandhorst/grpc-reflection-testing/proto"
)

//go:generate protoc -I vendor -I $GOPATH/src -I . --go_out=plugins=grpc:$GOPATH/src --govalidators_out=gogoimport=false:$GOPATH/src ./proto/example.proto

var (
	gRPCPort = flag.Int("grpc-port", 10000, "The gRPC server port")
)

var log grpclog.LoggerV2

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(log)
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *gRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(validator.StreamServerInterceptor()),
	)
	pbExample.RegisterUserServiceServer(s, srv{})

	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))
}

type srv struct{}

func (s srv) AddUser(ctx context.Context, req *pbExample.User) (*pbExample.Void, error) {
	return nil, nil
}
