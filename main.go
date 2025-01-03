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

// VarRef represents a variable reference
type VarRef struct {
	name string
}

// Fun represents a function definition
type Fun struct {
	param string
	body  AST
	arg   AST // The argument to apply when evaluating the function
}

// Add represents an addition operation with two values
type Add struct {
	left  AST
	right AST
}

// Sub represents a subtraction operation with two values
type Sub struct {
	left  AST
	right AST
}

// Eval evaluates an AST node and returns its result
func Eval(ast AST) int {
	return EvalWithEnv(ast, make(map[string]int))
}

// EvalWithEnv evaluates an AST node with a given variable environment
func EvalWithEnv(ast AST, env map[string]int) int {
	switch node := ast.(type) {
	case Add:
		return EvalWithEnv(node.left, env) + EvalWithEnv(node.right, env)
	case Sub:
		return EvalWithEnv(node.left, env) - EvalWithEnv(node.right, env)
	case Val:
		return node.value
	case VarRef:
		return env[node.name]
	case Fun:
		// Evaluate the argument and bind it to the parameter
		argValue := EvalWithEnv(node.arg, env)
		// Create a new environment with the parameter binding
		newEnv := make(map[string]int)
		for k, v := range env {
			newEnv[k] = v
		}
		newEnv[node.param] = argValue
		return EvalWithEnv(node.body, newEnv)
	default:
		return 0
	}
}

// Parse converts a string expression into an operation struct
func Parse(input string) AST {
	parts := strings.Split(input, " ")
	if len(parts) < 1 {
		return nil
	}

	switch parts[0] {
	case "ADD":
		if len(parts) < 3 {
			return nil
		}
		left := parseValue(parts[1])
		right := parseValue(parts[2])
		return Add{left, right}
	case "SUB":
		if len(parts) < 3 {
			return nil
		}
		left := parseValue(parts[1])
		right := parseValue(parts[2])
		return Sub{left, right}
	case "FUN":
		if len(parts) < 4 {
			return nil
		}
		param := parts[1]
		// The last part is the argument
		arg := parseValue(parts[len(parts)-1])
		// Parse the middle parts as the body expression
		bodyStr := strings.Join(parts[2:len(parts)-1], " ")
		body := Parse(bodyStr)
		return Fun{param, body, arg}
	default:
		return nil
	}
}

// parseValue converts a string to either a Val or VarRef
func parseValue(s string) AST {
	if val, err := strconv.Atoi(s); err == nil {
		return Val{val}
	}
	return VarRef{s}
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