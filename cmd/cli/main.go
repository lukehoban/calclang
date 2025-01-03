package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lukehoban/calclang/pkg/calculator"
)

func main() {
	fmt.Printf("CALCLANG CLI\n\n")

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

		expr := calculator.Parse(input)
		if expr == nil {
			fmt.Println("Invalid expression")
			continue
		}
		fmt.Println(calculator.Eval(expr))
	}
}