package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Adder interface {
	Add(ctx context.Context, x, y int) (int, error)
}

type adder struct {
	weaver.Implements[Adder]
}

func (a adder) Add(ctx context.Context, x, y int) (int, error) {
	a.Logger(ctx).Debug("Add", "x", x, "y", y)
	return x + y, nil
}
