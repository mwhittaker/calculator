package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ServiceWeaver/weaver"
)

//go:generate weaver generate ./...

type Server struct {
	weaver.Implements[weaver.Main]
	adder      weaver.Ref[Adder]
	multiplier weaver.Ref[Multiplier]
	lis        weaver.Listener `weaver:"calc"`
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
	product, err := s.multiplier.Get().Multiply(r.Context(), x, y)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, product)
}
