package calculator

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		want     Result
		wantErr  bool
	}{
		{
			name:  "simple addition",
			input: "ADD 5 3",
			want:  Result{Value: 8},
		},
		{
			name:  "simple subtraction",
			input: "SUB 10 4",
			want:  Result{Value: 6},
		},
		{
			name:     "invalid operation",
			input:    "MUL 5 3",
			want:     Result{Error: "invalid operation: MUL"},
			wantErr:  true,
		},
		{
			name:     "invalid format",
			input:    "ADD 5",
			want:     Result{Error: "invalid expression format: expected 3 parts, got 2"},
			wantErr:  true,
		},
		{
			name:     "invalid number",
			input:    "ADD 5 x",
			want:     Result{Error: "invalid right operand: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Calculate(tt.input)
			if tt.wantErr {
				if got.Error == "" {
					t.Errorf("Calculate() error = nil, wantErr %v", tt.wantErr)
				}
			} else {
				if got.Value != tt.want.Value {
					t.Errorf("Calculate() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}