package main

import (
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
		{"FUN x ADD x 1", Fun{"x", Add{Var{"x"}, Val{1}}}},
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
		input    string
		expected int
	}{
		{"ADD 1 2", 3},
		{"SUB 5 3", 2},
		{"FUN x ADD x 1 2", 3}, // Function that adds 1 to its argument, called with 2
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			expr := Parse(tc.input)
			env := make(Environment)
			res := Eval(expr, env)
			assert.Equal(t, tc.expected, res)
		})
	}
}
