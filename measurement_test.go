package main

import (
	"go/parser"
	"math"
	"testing"
)

// go test

func TestEval(t *testing.T) {
	m := &Measurement{
		units:  1.0,
		length: 2.0,
		width:  3.0,
		height: 4.0,
	}

	tests := []struct {
		expr     string
		expected float64
	}{
		{"a + b", 3.0},
		{"c * 2", 6.0},
		{"d ^ 2", 16.0},
		{"p + 1", math.Pi + 1},
		{"2.5 * (a + b)", 7.5},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			node, err := parser.ParseExpr(test.expr)
			if err != nil {
				t.Fatalf("failed to parse expression: %v", err)
			}
			result := m.eval(node)
			if result != test.expected {
				t.Errorf("for expression %s, expected %f, but got %f", test.expr, test.expected, result)
			}
		})
	}
}

// Rest of your code (Measurement struct and eval function) goes here
