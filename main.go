package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

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
	result := make(map[string]interface{})
	result["function"] = function
	result["left"] = a
	result["right"] = b
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
	result["fa"] = fa
	result["fb"] = fb
	result["kMax"] = k
	return result
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}

	t.ExecuteTemplate(w, "index", nil)
}

// Bisection Search Method Page Handler
func bisectionHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/bisection.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}

	if r.Method == "POST" {
		fmt.Println("its POST request and I'm handling it")
		// answer := make(map[string]string)
		function = r.FormValue("function")
		left, _ := strconv.ParseFloat(r.FormValue("leftEndPoint"), 64)
		right, _ := strconv.ParseFloat(r.FormValue("rightEndPoint"), 64)
		tol, _ := strconv.ParseFloat(r.FormValue("tolerance"), 64)
		kMax, _ := strconv.ParseInt(r.FormValue("kMax"), 0, 64)

		answer := bisection(left, right, tol, kMax)
		fmt.Println(answer)
		t.ExecuteTemplate(w, "bisection", answer)
	} else {
		answer := make(map[string]string)
		answer["function"] = ""
		answer["left"] = ""
		answer["right"] = ""
		answer["fa"] = ""
		answer["fb"] = ""
		answer["kMax"] = ""
		fmt.Println("it's a GET request and I'm handling it")
		t.ExecuteTemplate(w, "bisection", answer)
	}

	// t.ExecuteTemplate(w, "bisection", answer)
}

// Even Search Method Page Handler
func esmHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/esm.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}

	t.ExecuteTemplate(w, "esm", nil)
}

// Golden Search Method Page Handler
func goldenHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/golden.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}

	t.ExecuteTemplate(w, "golden", nil)
}

// Pocket Search Method Page Handler
func pocketHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/pocket.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}

	t.ExecuteTemplate(w, "pocket", nil)
}

func main() {
	fmt.Println("Listening on port 3000:")
	// Setting the path of templates
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// Route Handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bisection", bisectionHandler)
	http.HandleFunc("/esm", esmHandler)
	http.HandleFunc("/golden", goldenHandler)
	http.HandleFunc("/pocket", pocketHandler)

	http.ListenAndServe(":3000", nil)
}
