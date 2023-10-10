package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/mwhittaker/calculator/1/adder"
	"github.com/mwhittaker/calculator/1/calc"
	"github.com/mwhittaker/calculator/1/multiplier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr           = flag.String("addr", "localhost:8000", "Server address")
	adderAddr      = flag.String("adder", "localhost:8001", "Adder server address")
	multiplierAddr = flag.String("multiplier", "localhost:8002", "Multiplier server address")
)

func main() {
	flag.Parse()

	// Create a client to the adder service.
	adderConn, err := grpc.Dial(*adderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer adderConn.Close()
	adder := adder.NewAdderClient(adderConn)

	// Create a client to the multiplier service.
	multiplierConn, err := grpc.Dial(*multiplierAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer multiplierConn.Close()
	multiplier := multiplier.NewMultiplierClient(multiplierConn)

	// Serve HTTP traffic.
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Calculator server listening at %v\n", lis.Addr())
	if err := http.Serve(lis, calc.NewServer(adder, multiplier)); err != nil {
		panic(err)
	}
}
