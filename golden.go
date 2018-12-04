package main

import (
	"fmt"
	"math"

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

func goldenMin(a float64, b float64, tol float64, kMax int) {
	fmt.Println("\n\t\tGolden Search Method to find the minimum")
	r := (math.Sqrt(5.) - 1) / 2
	x1 := a + (1-r)*(b-a)
	f1 := f(x1)
	x2 := a + r*(b-a)
	f2 := f(x2)
	k := 0

	for k < kMax {
		k++
		if f1.(float64) > f2.(float64) {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + r*(b-a)
			f2 = f(x2)
		} else {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + (1-r)*(b-a)
			f1 = f(x1)
		}
	}

	fmt.Println("f1: ", f1)
	fmt.Println("x1: ", x1)
	fmt.Println("k: ", k)
	fmt.Println("abs(b-a): ", math.Abs(b-a))
}

func goldenMax(a float64, b float64, tol float64, kMax int) {
	fmt.Println("\n\n\t\tGolden Search Method to find the maximum")
	r := (math.Sqrt(5.) - 1) / 2
	x1 := a + (1-r)*(b-a)
	f1 := f(x1)
	x2 := a + r*(b-a)
	f2 := f(x2)
	k := 0

	for k < kMax {
		k++
		if f1.(float64) < f2.(float64) {
			a = x1
			x1 = x2
			f1 = f2
			x2 = a + r*(b-a)
			f2 = f(x2)
		} else {
			b = x2
			x2 = x1
			f2 = f1
			x1 = a + (1-r)*(b-a)
			f1 = f(x1)
		}
	}

	fmt.Println("f1: ", f1)
	fmt.Println("x1: ", x1)
	fmt.Println("k: ", k)
	fmt.Println("abs(b-a): ", math.Abs(b-a))
}

func main() {
	fmt.Print("Enter function: ")
	fmt.Scan(&function)
	var a float64
	var b float64
	var tol float64
	var kMax int
	fmt.Print("Enter a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter b: ")
	fmt.Scanln(&b)
	fmt.Print("Enter tol: ")
	fmt.Scanln(&tol)
	fmt.Print("Enter kMax: ")
	fmt.Scanln(&kMax)
	goldenMin(a, b, tol, kMax)
	goldenMax(a, b, tol, kMax)
}
