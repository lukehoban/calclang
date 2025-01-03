package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"ADD 1 2", Add{Val{1}, Val{2}}},
		{"SUB 1 2", Sub{Val{1}, Val{2}}},
		{"FUN x ADD x 1", Fun{"x", Add{Val{0}, Val{1}}}},
		{"FUN x ADD 1 x", Fun{"x", Add{Val{1}, Val{0}}}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			res := Parse(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestEval(t *testing.T) {
	tests := []struct {
		input    AST
		expected int
	}{
		{Add{Val{1}, Val{2}}, 3},
		{Sub{Val{1}, Val{2}}, -1},
		{Fun{"x", Add{Val{0}, Val{1}}}, 0}, // Function definition evaluates to 0
		{Call{Fun{"x", Add{Val{0}, Val{1}}}, Val{2}}, 3}, // Function call with x=2: 2+1=3
		{Call{Fun{"x", Add{Val{1}, Val{0}}}, Val{2}}, 3}, // Function call with x=2: 1+2=3
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			res := Eval(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
