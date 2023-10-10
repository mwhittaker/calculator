package calc

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mwhittaker/calculator/2/adder"
	"github.com/mwhittaker/calculator/2/multiplier"
)

type Server struct {
	adder      adder.AdderClient
	multiplier multiplier.MultiplierClient
}

func NewServer(adder adder.AdderClient, multiplier multiplier.MultiplierClient) *Server {
	return &Server{adder, multiplier}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var mux http.ServeMux
	mux.HandleFunc("/add", s.handleAdd)
	mux.HandleFunc("/multiply", s.handleMultiply)
	mux.ServeHTTP(w, r)
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
	summands := &adder.Summands{X: int64(x), Y: int64(y)}
	sum, err := s.adder.Add(r.Context(), summands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, sum.Sum)
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
