package main

import (
	"go/parser"
	"testing"
)

// go test -bench .

func BenchmarkEval(b *testing.B) {
	m := &Measurement{
		units:  1.0,
		length: 2.0,
		width:  3.0,
		height: 4.0,
	}

	expr := "a + b * c - d / 2.0" // Replace with an expression you want to benchmark

	node, err := parser.ParseExpr(expr)
	if err != nil {
		b.Fatalf("failed to parse expression: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m.eval(node)
	}
}
