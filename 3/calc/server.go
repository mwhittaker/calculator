package calc

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ServiceWeaver/weaver"
	"github.com/mwhittaker/calculator/3/multiplier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate weaver generate ./...

type config struct {
	MultiplierAddr string `toml:"multiplier_addr"`
}

type Server struct {
	weaver.Implements[weaver.Main]
	adder      weaver.Ref[Adder]
	multiplier multiplier.MultiplierClient
	weaver.WithConfig[config]
	lis weaver.Listener `weaver:"calc"`
}

func (s *Server) Init(context.Context) error {
	// Create a client to the multiplier service.
	conn, err := grpc.Dial(s.Config().MultiplierAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	s.multiplier = multiplier.NewMultiplierClient(conn)
	return nil
}

func Serve(ctx context.Context, s *Server) error {
	var mux http.ServeMux
	mux.HandleFunc("/add", s.handleAdd)
	mux.HandleFunc("/multiply", s.handleMultiply)
	s.Logger(ctx).Info("Calculator server listening", "addr", s.lis)
	return http.Serve(s.lis, &mux)
}

func (s *Server) handleAdd(w http.ResponseWriter, r *http.Request) {
	// Parse x and y from the URL.
	x, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid x: %v", err), http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid y: %v", err), http.StatusBadRequest)
		return
	}

	// Add x and y.
	sum, err := s.adder.Get().Add(r.Context(), x, y)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, sum)
}

func (s *Server) handleMultiply(w http.ResponseWriter, r *http.Request) {
	// Parse x and y from the URL.
	x, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid x: %v", err), http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid y: %v", err), http.StatusBadRequest)
		return
	}

	// Multiply x and y.
	multiplicands := &multiplier.Multiplicands{X: int64(x), Y: int64(y)}
	product, err := s.multiplier.Multiply(r.Context(), multiplicands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, product.Product)
}
