package main

import (
	"bufio"
	"fmt"
	"os"
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
func Parse(input string) AST {
	parts := strings.Split(input, " ")
	left, _ := strconv.Atoi(parts[1])
	right, _ := strconv.Atoi(parts[2])

	switch parts[0] {
	case "ADD":
		return Add{Val{left}, Val{right}}
	case "SUB":
		return Sub{Val{left}, Val{right}}
	default:
		return nil
	}
}

func main() {
	fmt.Printf("CALCLANG\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "" {
			continue
		}

		expr := Parse(input)
		if expr == nil {
			fmt.Println("Invalid expression")
			continue
		}
		fmt.Println(Eval(expr))
	}
}
