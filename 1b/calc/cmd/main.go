package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	app "github.com/mwhittaker/calculator/1b/calc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	calc              = flag.String("calc", "", "Run a calc server with the given address")
	adderAddress      = flag.String("adder_address", "", "Multiplier server address")
	multiplierAddress = flag.String("multiplier_address", "", "Adder server address")
	adder             = flag.String("adder", "", "Run an adder server with the given address")
	multiplier        = flag.String("multiplier", "", "Run a multiplier server with the given address")
)

func main() {
	flag.Parse()
	if *calc == "" && *adder == "" && *multiplier == "" {
		// If no flags are given, run everything locally.
		*calc = "localhost:8000"
		*adder = "localhost:8001"
		*multiplier = "localhost:8002"
		*adderAddress = *adder
		*multiplierAddress = *multiplier
	}

	var group errgroup.Group

	if *adder != "" {
		// Run the adder gRPC server.
		lis, err := net.Listen("tcp", *adder)
		if err != nil {
			panic(err)
		}
		s := grpc.NewServer()
		app.RegisterAdderServer(s, &app.Adder{})
		fmt.Printf("Adder server listening at %v\n", lis.Addr())
		group.Go(func() error {
			return s.Serve(lis)
		})
	}

	if *multiplier != "" {
		// Run the multiplier gRPC server.
		lis, err := net.Listen("tcp", *multiplier)
		if err != nil {
			panic(err)
		}
		s := grpc.NewServer()
		app.RegisterMultiplierServer(s, &app.Multiplier{})
		fmt.Printf("Multiplier server listening at %v\n", lis.Addr())
		group.Go(func() error {
			return s.Serve(lis)
		})
	}

	if *calc != "" {
		// Run the calc HTTP server.

		// Create a client to the adder service.
		adderConn, err := grpc.Dial(*adderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer adderConn.Close()
		adder := app.NewAdderClient(adderConn)

		// Create a client to the multiplier service.
		multiplierConn, err := grpc.Dial(*multiplierAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer multiplierConn.Close()
		multiplier := app.NewMultiplierClient(multiplierConn)

		// Serve HTTP traffic.
		lis, err := net.Listen("tcp", *calc)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Calculator server listening at %v\n", lis.Addr())
		group.Go(func() error {
			return http.Serve(lis, app.NewCalc(adder, multiplier))
		})
	}

	if err := group.Wait(); err != nil {
		panic(err)
	}
}
