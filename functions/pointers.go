package main

import (
	"fmt"
	"reflect"
)

func main() {

	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

	myInt := 4
	myIntPointer := &myInt
	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(&myIntPointer), "Pointer to a Pointer")
	fmt.Println(&myIntPointer, "Pointer address of pointer")
	fmt.Println(myIntPointer, "Pointer address of value")
	fmt.Println(*myIntPointer, "Pointer value")
	//assign a new value
	*myIntPointer = 8
	fmt.Println(*myIntPointer, "changing value with pointer")
	fmt.Println(myInt, "Print values directly")

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(&myFloat)
	fmt.Println(reflect.TypeOf(&myFloatPointer), "Pointer to a Pointer")
	fmt.Println(&myFloatPointer, "Pointer address of pointer")
	fmt.Println(myFloatPointer, "Pointer address of value")
	fmt.Println(*myFloatPointer, "Pointer value")

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(reflect.TypeOf(&myBool))
	fmt.Println(&myBool)
	fmt.Println(reflect.TypeOf(&myBoolPointer), "Pointer to a Pointer")
	fmt.Println(&myBoolPointer, "Pointer address of pointer")
	fmt.Println(myBoolPointer, "Pointer address of value")
	fmt.Println(*myBoolPointer, "Pointer value")

	myString := "test"
	myStringPointer := &myString
	fmt.Println(reflect.TypeOf(&myString))
	fmt.Println(&myString)
	fmt.Println(reflect.TypeOf(&myStringPointer), "Pointer to a Pointer")
	fmt.Println(&myStringPointer, "Pointer address of pointer")
	fmt.Println(myStringPointer, "Pointer address of value")
	fmt.Println(*myStringPointer, "Pointer value")

	var myFloatPointers *float64 = createPointer()
	fmt.Println(*myFloatPointers)
	fmt.Println(myFloatPointers)

	var myBools bool = true
	printPointer(&myBools)

	amount = 6
	double(&amount)
	fmt.Println(amount)

	truth := true
	negate(&truth)
	fmt.Println(truth)

	lies := false
	negate(&lies)
	fmt.Println(lies)
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(number *int) {
	*number *= 2
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
