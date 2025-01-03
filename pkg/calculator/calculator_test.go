package calculator

import (
	"testing"

	"gotest.tools/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
	}{
		{"add operation", "ADD 1 2", Add{Val{1}, Val{2}}},
		{"sub operation", "SUB 1 2", Sub{Val{1}, Val{2}}},
		{"invalid operation", "MUL 1 2", nil},
		{"invalid format", "ADD 1", nil},
		{"invalid numbers", "ADD a b", nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := Parse(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestEval(t *testing.T) {
	tests := []struct {
		name     string
		input    AST
		expected int
	}{
		{"add operation", Add{Val{1}, Val{2}}, 3},
		{"sub operation", Sub{Val{5}, Val{3}}, 2},
		{"single value", Val{42}, 42},
		{"nil input", nil, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := Eval(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}