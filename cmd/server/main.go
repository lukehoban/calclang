package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lukehoban/calclang/pkg/calculator"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := calculator.Calculate(req.Expression)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	port := ":8080"
	http.HandleFunc("/calculate", calculateHandler)
	
	fmt.Printf("CALCLANG Server starting on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}