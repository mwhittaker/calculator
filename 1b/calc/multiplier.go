package calc

import (
	"context"
	"fmt"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative multiplier.proto

type Multiplier struct {
	UnimplementedMultiplierServer
}

func (*Multiplier) Multiply(ctx context.Context, in *Multiplicands) (*Product, error) {
	fmt.Printf("Multiply(%d, %d)\n", in.X, in.Y)
	return &Product{Product: in.X * in.Y}, nil
}
