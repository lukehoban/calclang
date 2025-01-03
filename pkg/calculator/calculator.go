package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

// AST represents an abstract syntax tree node
type AST interface{}

// Val represents a numeric value
type Val struct {
	value int
}

// Add represents an addition operation with two values
type Add struct {
	left  Val
	right Val
}

// Sub represents a subtraction operation with two values
type Sub struct {
	left  Val
	right Val
}

// Result represents the result of a calculation
type Result struct {
	Value int    `json:"value"`
	Error string `json:"error,omitempty"`
}

// Eval evaluates an AST node and returns its result
func Eval(ast AST) int {
	switch node := ast.(type) {
	case Add:
		return node.left.value + node.right.value
	case Sub:
		return node.left.value - node.right.value
	case Val:
		return node.value
	default:
		return 0
	}
}

// Parse converts a string expression into an operation struct
func Parse(input string) (AST, error) {
	parts := strings.Split(strings.TrimSpace(input), " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid expression format: expected 3 parts, got %d", len(parts))
	}

	left, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid left operand: %v", err)
	}

	right, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid right operand: %v", err)
	}

	switch parts[0] {
	case "ADD":
		return Add{Val{left}, Val{right}}, nil
	case "SUB":
		return Sub{Val{left}, Val{right}}, nil
	default:
		return nil, fmt.Errorf("invalid operation: %s", parts[0])
	}
}

// Calculate processes an input string and returns a Result
func Calculate(input string) Result {
	expr, err := Parse(input)
	if err != nil {
		return Result{Error: err.Error()}
	}
	return Result{Value: Eval(expr)}
}