package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"slices"

	"github.com/BurntSushi/toml"
	app "github.com/mwhittaker/calculator/1b/calc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type node struct {
	HTTPAddress string   `toml:"http_address"`
	GRPCAddress string   `toml:"grpc_address"`
	Services    []string `toml:"services"`
}

type topology struct {
	Nodes map[string]node `toml:"nodes"`
}

func main() {
	flag.Parse()

	var topo topology
	var name string
	switch flag.NArg() {
	case 0:
		// NOTE(mwhittaker): Beaver can help run everything locally by default.
		topo = topology{
			Nodes: map[string]node{
				"main": {
					HTTPAddress: "localhost:8000",
					Services:    []string{"adder", "multiplier", "calc"},
				},
			},
		}
		name = "main"
	case 2:
		if _, err := toml.DecodeFile(flag.Arg(0), &topo); err != nil {
			panic(err)
		}
		name = flag.Arg(1)
	default:
		fmt.Fprintln(os.Stderr, "usage: cmd <topo file> <node>")
		os.Exit(1)
	}

	// NOTE(mwhittaker): Beaver can help run against a local gRPC service when
	// possible and a remote service otherwise.
	// NOTE(mwhittaker): Beaver can help make a gRPC client that load balances
	// across a set of servers.
	var adder app.AdderServer
	var multiplier app.MultiplierServer
	getAdder := func() app.AdderClient {
		if adder != nil {
			return app.NewAdderClient(app.Local(adder))
		}
		for _, node := range topo.Nodes {
			if slices.Contains(node.Services, "adder") {
				conn, err := grpc.Dial(node.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					panic(err)
				}
				return app.NewAdderClient(conn)
			}
		}
		panic("adder not found")
	}
	getMultiplier := func() app.MultiplierClient {
		if multiplier != nil {
			return app.NewMultiplierClient(app.Local(multiplier))
		}
		for _, node := range topo.Nodes {
			if slices.Contains(node.Services, "multiplier") {
				conn, err := grpc.Dial(node.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					panic(err)
				}
				return app.NewMultiplierClient(conn)
			}
		}
		panic("multiplier not found")
	}

	// NOTE(mwhittaker): You have to initialize services in topological order.
	// Maybe beaver can help declare the gRPC dependency graph and ensure
	// things are run in topological order. This seems similar to
	// https://github.com/google/wire to me?
	n := topo.Nodes[name]
	grpcServer := grpc.NewServer()
	var group errgroup.Group
	// NOTE(mwhittaker): beaver can automate this?
	if slices.Contains(n.Services, "adder") {
		adder = &app.Adder{}
		app.RegisterAdderServer(grpcServer, adder)
	}
	if slices.Contains(n.Services, "multiplier") {
		multiplier = &app.Multiplier{}
		app.RegisterMultiplierServer(grpcServer, multiplier)
	}
	if slices.Contains(n.Services, "calc") {
		adder := getAdder()
		multiplier := getMultiplier()
		// Serve HTTP traffic.
		lis, err := net.Listen("tcp", n.HTTPAddress)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Calculator server listening at %v\n", lis.Addr())
		group.Go(func() error {
			return http.Serve(lis, app.NewCalc(adder, multiplier))
		})
	}

	if n.GRPCAddress != "" {
		lis, err := net.Listen("tcp", n.GRPCAddress)
		if err != nil {
			panic(err)
		}
		fmt.Printf("GRPC servers listening at %v\n", lis.Addr())
		group.Go(func() error {
			return grpcServer.Serve(lis)
		})
	}

	if err := group.Wait(); err != nil {
		panic(err)
	}
}
