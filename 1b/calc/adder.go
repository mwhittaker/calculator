package calc

import (
	"context"
	"fmt"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative adder.proto

type Adder struct {
	UnimplementedAdderServer
}

func (*Adder) Add(ctx context.Context, in *Summands) (*Sum, error) {
	fmt.Printf("Add(%d, %d)\n", in.X, in.Y)
	return &Sum{Sum: in.X + in.Y}, nil
}
