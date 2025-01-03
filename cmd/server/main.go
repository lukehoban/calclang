package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lukehoban/calclang/pkg/calculator"
)

type CalculationRequest struct {
	Expression string `json:"expression"`
}

type CalculationResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	expr := calculator.Parse(req.Expression)
	if expr == nil {
		json.NewEncoder(w).Encode(CalculationResponse{
			Error: "Invalid expression",
		})
		return
	}

	result := calculator.Eval(expr)
	json.NewEncoder(w).Encode(CalculationResponse{
		Result: result,
	})
}

func main() {
	fmt.Println("CALCLANG Server starting on :8080")
	
	http.HandleFunc("/calculate", calculateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}