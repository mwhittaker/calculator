package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Multiplier interface {
	Multiply(ctx context.Context, x, y int) (int, error)
}

type multiplier struct {
	weaver.Implements[Multiplier]
}

func (m multiplier) Multiply(ctx context.Context, x, y int) (int, error) {
	m.Logger(ctx).Debug("Multiply", "x", x, "y", y)
	return x * y, nil
}
