package main

import (
	"fmt"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func englishHander(writer http.ResponseWriter, request *http.Request) {
	write(writer, "hello, Web")
}

func frenchHander(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut, Web")
}

func malayalamHander(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaskaram, Web")
}
func d(writer http.ResponseWriter, request *http.Request) {
	write(writer, "z")
}

func e(writer http.ResponseWriter, request *http.Request) {
	write(writer, "x")
}

func f(writer http.ResponseWriter, request *http.Request) {
	write(writer, "y")
}

func sayHi() {
	fmt.Println("hi")
}

func sayBye() {
	fmt.Println("bye")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func multiply(a int, b int) float64 {
	return float64(a) * float64(b)
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func doMath(passedFunction func(int, int) float64) {
	result := passedFunction(10, 2)
	fmt.Println(result)
}

func calledFunction(passedFunction func()) {
	passedFunction()
}

func calledTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("This sentance is", false)
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {

	calledFunction(functionA)
	calledTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)

	doMath(divide)
	doMath(multiply)
	twice(sayHi)
	twice(sayBye)

	//variable holds no parameters
	var greeterFunction func()
	//variable will hold a function with two int parameters and a float64 return value
	var mathFunction func(int, int) float64
	greeterFunction = sayHi
	mathFunction = divide

	greeterFunction()
	fmt.Println(mathFunction(5, 2))

	http.HandleFunc("/hello", englishHander)
	http.HandleFunc("/salut", frenchHander)
	http.HandleFunc("/namaskaram", malayalamHander)
	http.HandleFunc("/a", f)
	http.HandleFunc("/b", d)
	http.HandleFunc("/c", e)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
