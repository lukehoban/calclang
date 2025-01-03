package main

import (
	"fmt"
	"github.com/lukehoban/calclang"
)

func main() {
	input := "FUN x ADD x 1 2"
	expr := calclang.Parse(input)
	env := make(calclang.Environment)
	result := calclang.Eval(expr, env)
	fmt.Printf("Result: %d\n", result)
}