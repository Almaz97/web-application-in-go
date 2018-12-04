package main

import (
	"fmt"

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

func esmMin(x0 float64, h float64, kMax int) {
	fmt.Println("\n\t\tEven Search Method to find the minimum")
	var x1 float64
	k := 0
	yf0 := f(x0)
	yf1 := f(x0 + h)
	for k < kMax {
		k++
		x1 = x0 + h
		if yf1.(float64) >= yf0.(float64) {
			x1 = x0
			yf1 = yf0
		} else {
			x0 = x1
			yf0 = yf1
			x1 = x0 + h
			yf1 = f(x1)
		}
	}

	fmt.Println("x1:", x1)
	fmt.Println("yf1:", yf1)
	fmt.Println("k:", k)
	fmt.Println("h:", h)
}

func esmMax(x0 float64, h float64, kMax int) {
	fmt.Println("\n\n\t\tEven Search Method to find the maximum")
	var x1 float64
	k := 0
	yf0 := f(x0)
	yf1 := f(x0 + h)
	for k < kMax {
		k++
		x1 = x0 + h
		if yf1.(float64) <= yf0.(float64) {
			x1 = x0
			yf1 = yf0
		} else {
			x0 = x1
			yf0 = yf1
			x1 = x0 + h
			yf1 = f(x1)
		}
	}

	fmt.Println("x1:", x1)
	fmt.Println("yf1:", yf1)
	fmt.Println("k:", k)
	fmt.Println("h:", h)
}

func main() {
	fmt.Print("Enter function: ")
	fmt.Scan(&function)
	var x0 float64
	var h float64
	var kMax int
	fmt.Print("Enter x0: ")
	fmt.Scanln(&x0)
	fmt.Print("Enter h: ")
	fmt.Scanln(&h)
	fmt.Print("Enter kMax: ")
	fmt.Scanln(&kMax)

	esmMin(x0, h, kMax)
	esmMax(x0, h, kMax)
}
