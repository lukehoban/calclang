package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AST interface{}

type Val struct {
	value int
}

type VarRef struct {
	name string
}

type Fun struct {
	param string
	body  AST
	arg   AST
}

type Add struct {
	left  AST
	right AST
}

type Sub struct {
	left  AST
	right AST
}

func Eval(ast AST) int {
	return EvalWithEnv(ast, make(map[string]int))
}

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
		argValue := EvalWithEnv(node.arg, env)
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
		arg := parseValue(parts[len(parts)-1])
		bodyStr := strings.Join(parts[2:len(parts)-1], " ")
		body := Parse(bodyStr)
		return Fun{param, body, arg}
	default:
		return nil
	}
}

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