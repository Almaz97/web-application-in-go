package main

import (
	"fmt"
	"math"
	"os"

	"github.com/Knetic/govaluate"
)

var function string

func f(x float64) interface{} {
	expression, _ := govaluate.NewEvaluableExpression(function)
	parameters := make(map[string]interface{}, 8)
	parameters["x"] = x
	result, _ := expression.Evaluate(parameters)

	return result
}

func bisection(a float64, b float64, tol float64, kMax int64) map[string]interface{} {
	fmt.Println("\n\n\t\tBisection Search Method")
	var k int64
	fa := f(a)
	fb := f(b)

	if math.Signbit(fa.(float64)) == math.Signbit(fb.(float64)) {
		fmt.Println("Sign of f(a) and f(b) must be opposite! Check endpoints of the internal [a,b]")
		os.Exit(1)
	}

	k = 0

	for b-a > tol && k < kMax {
		m := a + (b-a)/2
		fa = f(a)
		fb = f(m)
		if math.Signbit(fa.(float64)) == math.Signbit(fb.(float64)) {
			a = m
		} else {
			b = m
		}
		k++
	}
	result := make(map[string]interface{})
	result["left"] = a
	result["right"] = b
	result["fa"] = fa
	result["fb"] = fb
	result["kMax"] = k
	return result
}

// func main() {

// 	fmt.Print("Enter function: ")
// 	fmt.Scan(&function)

// 	var a float64
// 	var b float64
// 	var tol float64
// 	var kMax int
// 	fmt.Print("Enter a: ")
// 	fmt.Scanln(&a)
// 	fmt.Print("Enter b: ")
// 	fmt.Scanln(&b)
// 	fmt.Print("Enter tol: ")
// 	fmt.Scanln(&tol)
// 	fmt.Print("Enter kMax: ")
// 	fmt.Scanln(&kMax)

// 	bisection(a, b, tol, kMax)

// }
