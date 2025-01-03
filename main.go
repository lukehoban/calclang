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
	left  AST
	right AST
}

// Sub represents a subtraction operation with two values
type Sub struct {
	left  AST
	right AST
}

// Fun represents a function definition
type Fun struct {
	param string
	body  AST
}

// Call represents a function call
type Call struct {
	fun  AST
	arg  AST
}

// Var represents a variable reference
type Var struct {
	name string
}

// Environment holds variable bindings
type Environment map[string]AST

// Eval evaluates an AST node and returns its result
func Eval(ast AST, env Environment) int {
	switch node := ast.(type) {
	case Add:
		return Eval(node.left, env) + Eval(node.right, env)
	case Sub:
		return Eval(node.left, env) - Eval(node.right, env)
	case Val:
		return node.value
	case Fun:
		return 0 // Functions don't evaluate to a value directly
	case Call:
		if fun, ok := node.fun.(Fun); ok {
			// Create new environment with argument bound to parameter
			newEnv := make(Environment)
			for k, v := range env {
				newEnv[k] = v
			}
			newEnv[fun.param] = node.arg
			return Eval(fun.body, newEnv)
		}
		return 0
	case Var:
		if val, ok := env[node.name]; ok {
			return Eval(val, env)
		}
		return 0
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
		if len(parts) != 3 {
			return nil
		}
		left := parseValue(parts[1])
		right := parseValue(parts[2])
		return Add{left, right}
	case "SUB":
		if len(parts) != 3 {
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
		// Check if this is a function call (has an argument at the end)
		if len(parts) > 4 {
			// The last part is the argument
			argStr := parts[len(parts)-1]
			// Everything in between is the function body
			bodyStr := strings.Join(parts[2:len(parts)-1], " ")
			body := Parse(bodyStr)
			arg := parseValue(argStr)
			return Call{Fun{param, body}, arg}
		}
		// Just a function definition
		bodyStr := strings.Join(parts[2:], " ")
		body := Parse(bodyStr)
		return Fun{param, body}
	default:
		// Try to parse as a number
		if val := parseValue(parts[0]); val != nil {
			return val
		}
		// Must be a variable
		return Var{parts[0]}
	}
}

// parseValue tries to parse a string as a number or returns it as a variable reference
func parseValue(s string) AST {
	if n, err := strconv.Atoi(s); err == nil {
		return Val{n}
	}
	return Var{s}
}

func main() {
	fmt.Printf("CALCLANG\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	env := make(Environment)

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
		fmt.Println(Eval(expr, env))
	}
}