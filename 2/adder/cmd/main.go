package main

import (
	"context"
	"flag"

	"github.com/ServiceWeaver/weaver"
	"github.com/mwhittaker/calculator/2/adder"
)

var addr = flag.String("addr", "localhost:8001", "Server address")

func main() {
	flag.Parse()
	if err := weaver.Run(context.Background(), adder.Serve); err != nil {
		panic(err)
	}
}
