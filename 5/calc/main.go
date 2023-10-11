package main

import (
	"context"
	"flag"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	flag.Parse()
	if err := weaver.Run(context.Background(), Serve); err != nil {
		panic(err)
	}
}
