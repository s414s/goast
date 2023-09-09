package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"strconv"
	"strings"
)

type Measurement struct {
	units   float64
	length  float64
	width   float64
	height  float64
	formula string
	comment string
	typeM   string
}

func (m *Measurement) TotalAmount() float64 {
	// Aqui se hacen todos los calculos en el futuro
	exp, err := parser.ParseExpr(m.formula)
	if err != nil {
		fmt.Println("Error in expression")
		return 0
	}
	return m.eval(exp)
}

// TODO - hacer que se devuelvan errores en caso de tener caracteres no permitidos
func (m *Measurement) eval(node ast.Expr) float64 {
	switch expr := node.(type) {
	case *ast.BinaryExpr:
		left := m.eval(expr.X)
		right := m.eval(expr.Y)
		switch expr.Op.String() {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		case "^":
			return math.Pow(left, right)
		}
	case *ast.ParenExpr:
		return m.eval(expr.X)
	case *ast.BasicLit:
		if expr.Kind == token.INT {
			val, err := strconv.ParseFloat(expr.Value, 64)
			if err != nil {
				fmt.Print("Error con el valor")
				return 0
			}
			return val
		} else if expr.Kind == token.FLOAT {
			val, err := strconv.ParseFloat(expr.Value, 64)
			if err != nil {
				fmt.Print("Error con el valor")
				return 0
			}
			return val
		}
	case *ast.Ident:
		switch strings.ToLower(expr.Name) {
		case "a":
			return m.units
		case "b":
			return m.length
		case "c":
			return m.width
		case "d":
			return m.height
		case "p":
			return math.Pi
		default:
			return 0
		}
	}

	return 0
}

func calculate(node ast.Expr) int {
	switch expr := node.(type) {
	case *ast.BinaryExpr:
		left := calculate(expr.X)
		right := calculate(expr.Y)
		switch expr.Op {
		case token.ADD:
			return left + right
		case token.SUB:
			return left - right
		case token.MUL:
			return left * right
		case token.QUO:
			return left / right
		}
	case *ast.ParenExpr:
		return calculate(expr.X)
	case *ast.BasicLit:
		if expr.Kind == token.INT {
			val, error := strconv.Atoi(expr.Value)
			if error != nil {
				fmt.Print("Error con el valor")
				break
			}
			return val
		}
	}

	return 0
}

func Eval(exp ast.Expr) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.ParenExpr:
		return Eval(exp.X)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, err := strconv.ParseFloat(exp.Value, 64)
			if err != nil {
				fmt.Print("Error when parsing int")
				return 0
			}
			return i
		case token.FLOAT:
			i, err := strconv.ParseFloat(exp.Value, 64)
			if err != nil {
				fmt.Print("Error when parsing float")
				return 0
			}
			return i
		}
	case *ast.Ident:
		switch exp.Name {
		case "a":
			// TODO - sustituir por el valor en cuestion
			return 1
		case "b":
			// TODO - sustituir por el valor en cuestion
			return 2
		case "c":
			// TODO - sustituir por el valor en cuestion
			return 3
		case "d":
			// TODO - sustituir por el valor en cuestion
			return 4
		case "p":
			// TODO - sustituir por el valor en cuestion
			return math.Pi
		default:
			return 0
		}
	}

	fmt.Print("Antes del return 0 del Eval \n")
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) float64 {
	left := Eval(exp.X)
	right := Eval(exp.Y)

	fmt.Print("Binay expression detectada")
	// fmt.Print(exp.Op)
	fmt.Print(exp.Op.String())
	// Dado que no existe el tipo Token Power ni Square Root, voy a tener que usar los strings

	// switch exp.Op {
	// case token.ADD:
	// 	return left + right
	// case token.SUB:
	// 	return left - right
	// case token.MUL:
	// 	return left * right
	// case token.QUO:
	// 	return left / right
	// }

	switch exp.Op.String() {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	case "^":
		return math.Pow(float64(left), float64(right))
	}

	return 0
}

func main() {
	// We just have to use this to make the printer believe that we got our Go source code from a file
	// fs := token.NewFileSet()
	// exp, err := parser.ParseExpr("(3.4 - a) * 5")
	// if err != nil {
	// 	fmt.Println("Error in expression")
	// 	return
	// }
	// ast.Print(fs, exp)

	// fmt.Println("Expression clean")
	// printer.Fprint(os.Stdout, token.NewFileSet(), exp)
	// fmt.Printf("\n")

	// fmt.Println("===========Result================")
	// fmt.Printf("Resultado FINAL %f\n", Eval(exp))

	// fmt.Println("===========Con el calculate================")
	// result := calculate(exp)
	// fmt.Println("Result:", result)

	measurement := Measurement{
		comment: "sin comentario",
		units:   1.0,
		length:  2.0,
		width:   3.0,
		height:  4.0,
		formula: "(a + b) * p",
	}
	fmt.Printf("%f", measurement.TotalAmount())
}
