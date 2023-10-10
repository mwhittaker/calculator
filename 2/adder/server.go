package adder

import (
	"context"

	"github.com/ServiceWeaver/weaver"
	grpc "google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative adder.proto
//go:generate weaver generate ./...

type Server struct {
	weaver.Implements[weaver.Main]
	adder weaver.Ref[Adder]
	lis   weaver.Listener `weaver:"adder"`
}

func Serve(ctx context.Context, server *Server) error {
	s := grpc.NewServer()
	RegisterAdderServer(s, &grpcServer{})
	server.Logger(ctx).Info("Adder server listening", "addr", server.lis)
	return s.Serve(server.lis)
}

type grpcServer struct {
	UnimplementedAdderServer
	adder Adder
}

func (s *grpcServer) Add(ctx context.Context, in *Summands) (*Sum, error) {
	sum, err := s.adder.Add(ctx, int(in.X), int(in.Y))
	return &Sum{Sum: int64(sum)}, err
}
