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

// Fun represents a function definition
type Fun struct {
	param string
	body  AST
}

// Call represents a function call
type Call struct {
	fun  Fun
	arg  Val
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
	case Fun:
		// Functions evaluate to 0 until they are called
		return 0
	case Call:
		// When evaluating a function call, we need to substitute the argument value
		// into the function body and evaluate the result
		switch body := node.fun.body.(type) {
		case Add:
			left := body.left.value
			right := body.right.value
			// Replace parameter references with the argument value
			if body.left.value == 0 {
				left = node.arg.value
			}
			if body.right.value == 0 {
				right = node.arg.value
			}
			return left + right
		}
		return 0
	default:
		return 0
	}
}

// Parse converts a string expression into an operation struct
func Parse(input string) AST {
	parts := strings.Split(input, " ")
	
	switch parts[0] {
	case "ADD":
		left, _ := strconv.Atoi(parts[1])
		right, _ := strconv.Atoi(parts[2])
		return Add{Val{left}, Val{right}}
	case "SUB":
		left, _ := strconv.Atoi(parts[1])
		right, _ := strconv.Atoi(parts[2])
		return Sub{Val{left}, Val{right}}
	case "FUN":
		if len(parts) < 4 {
			return nil
		}
		param := parts[1]
		// Parse the body expression
		bodyParts := parts[2:]
		if bodyParts[0] == "ADD" {
			// Replace parameter references with the actual value
			if bodyParts[1] == param {
				bodyParts[1] = "0" // Placeholder, will be replaced during evaluation
			}
			if bodyParts[2] == param {
				bodyParts[2] = "0" // Placeholder, will be replaced during evaluation
			}
			left, _ := strconv.Atoi(bodyParts[1])
			right, _ := strconv.Atoi(bodyParts[2])
			return Fun{param, Add{Val{left}, Val{right}}}
		}
	default:
		return nil
	}
	return nil
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
