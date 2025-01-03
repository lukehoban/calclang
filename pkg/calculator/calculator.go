package calculator

import (
	"strconv"
	"strings"
)

// AST represents an abstract syntax tree node
type AST interface{}

// Val represents a numeric value
type Val struct {
	Value int
}

// Add represents an addition operation with two values
type Add struct {
	Left  Val
	Right Val
}

// Sub represents a subtraction operation with two values
type Sub struct {
	Left  Val
	Right Val
}

// Eval evaluates an AST node and returns its result
func Eval(ast AST) int {
	switch node := ast.(type) {
	case Add:
		return node.Left.Value + node.Right.Value
	case Sub:
		return node.Left.Value - node.Right.Value
	case Val:
		return node.Value
	default:
		return 0
	}
}

// Parse converts a string expression into an operation struct
func Parse(input string) AST {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return nil
	}
	
	left, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil
	}
	
	right, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil
	}

	switch parts[0] {
	case "ADD":
		return Add{Val{left}, Val{right}}
	case "SUB":
		return Sub{Val{left}, Val{right}}
	default:
		return nil
	}
}