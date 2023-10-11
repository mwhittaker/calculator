package main

import (
	"context"
	"flag"

	"github.com/ServiceWeaver/weaver"
	"github.com/mwhittaker/calculator/3b/calc"
)

func main() {
	flag.Parse()
	if err := weaver.Run(context.Background(), calc.Serve); err != nil {
		panic(err)
	}
}
