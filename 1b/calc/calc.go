package calc

import (
	"fmt"
	"net/http"
	"strconv"
)

type Calc struct {
	adder      AdderClient
	multiplier MultiplierClient
}

func NewCalc(adder AdderClient, multiplier MultiplierClient) *Calc {
	return &Calc{adder, multiplier}
}

func (s *Calc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var mux http.ServeMux
	mux.HandleFunc("/add", s.handleAdd)
	mux.HandleFunc("/multiply", s.handleMultiply)
	mux.ServeHTTP(w, r)
}

func (s *Calc) handleAdd(w http.ResponseWriter, r *http.Request) {
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
	summands := &Summands{X: int64(x), Y: int64(y)}
	sum, err := s.adder.Add(r.Context(), summands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, sum.Sum)
}

func (s *Calc) handleMultiply(w http.ResponseWriter, r *http.Request) {
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
	multiplicands := &Multiplicands{X: int64(x), Y: int64(y)}
	product, err := s.multiplier.Multiply(r.Context(), multiplicands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, product.Product)
}
