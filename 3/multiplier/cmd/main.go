package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/mwhittaker/calculator/3/multiplier"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:8002", "Server address")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	multiplier.RegisterMultiplierServer(s, &multiplier.Server{})
	fmt.Printf("Multiplier server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
