package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected AST
	}{
		{"ADD 1 2", Add{Val{1}, Val{2}}},
		{"SUB 1 2", Sub{Val{1}, Val{2}}},
		{"ADD x 2", Add{VarRef{"x"}, Val{2}}},
		{"FUN x ADD x 1 2", Fun{"x", Add{VarRef{"x"}, Val{1}}, Val{2}}},
		{"FUN y SUB y 5 3", Fun{"y", Sub{VarRef{"y"}, Val{5}}, Val{3}}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			res := Parse(tc.input)
			if res != tc.expected {
				t.Errorf("Parse(%q) = %v, want %v", tc.input, res, tc.expected)
			}
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
		{"FUN x ADD x 1 2", 3},  // x=2, so 2+1=3
		{"FUN y SUB y 5 3", -2}, // y=3, so 3-5=-2
		{"FUN z ADD z z 4", 8},  // z=4, so 4+4=8
		{"FUN a SUB 10 a 5", 5}, // a=5, so 10-5=5
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			expr := Parse(tc.input)
			result := Eval(expr)
			if result != tc.expected {
				t.Errorf("Eval(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestEvalWithEnv(t *testing.T) {
	tests := []struct {
		input    string
		env      map[string]int
		expected int
	}{
		{"ADD x y", map[string]int{"x": 1, "y": 2}, 3},
		{"SUB a b", map[string]int{"a": 5, "b": 3}, 2},
		{"FUN x ADD x y 2", map[string]int{"y": 3}, 5}, // x=2, y=3, so 2+3=5
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			expr := Parse(tc.input)
			result := EvalWithEnv(expr, tc.env)
			if result != tc.expected {
				t.Errorf("EvalWithEnv(%q, %v) = %v, want %v", tc.input, tc.env, result, tc.expected)
			}
		})
	}
}

func TestInvalidInputs(t *testing.T) {
	tests := []string{
		"",
		"ADD",
		"ADD 1",
		"SUB",
		"SUB 1",
		"FUN",
		"FUN x",
		"FUN x ADD",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			res := Parse(input)
			if res != nil {
				t.Errorf("Parse(%q) = %v, want nil", input, res)
			}
		})
	}
}