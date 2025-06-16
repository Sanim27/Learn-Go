package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add", add).Methods("GET")
	router.HandleFunc("/subtract", subtract).Methods("GET")
	router.HandleFunc("/multiply", multiply).Methods("GET")
	router.HandleFunc("/divide", divide).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func getOperands(r *http.Request) (float64, float64, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err1 := strconv.ParseFloat(aStr, 64)
	b, err2 := strconv.ParseFloat(bStr, 64)

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("invalid operands")
	}
	return a, b, nil
}

func writeJSON(w http.ResponseWriter, res Result) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func add(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		writeJSON(w, Result{Error: err.Error()})
		return
	}
	writeJSON(w, Result{Result: a + b})
}

func subtract(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		writeJSON(w, Result{Error: err.Error()})
		return
	}
	writeJSON(w, Result{Result: a - b})
}

func multiply(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		writeJSON(w, Result{Error: err.Error()})
		return
	}
	writeJSON(w, Result{Result: a * b})
}

func divide(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		writeJSON(w, Result{Error: err.Error()})
		return
	}
	if b == 0 {
		writeJSON(w, Result{Error: "Cannot divide by 0"})
		return
	}
	writeJSON(w, Result{Result: a / b})
}
