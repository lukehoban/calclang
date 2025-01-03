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

		result := calculator.Calculate(input)
		if result.Error != "" {
			fmt.Printf("Error: %s\n", result.Error)
			continue
		}
		fmt.Println(result.Value)
	}
}