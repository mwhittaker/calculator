package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/mwhittaker/calculator/1/adder"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:8001", "Server address")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	adder.RegisterAdderServer(s, &adder.Server{})
	fmt.Printf("Adder server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
