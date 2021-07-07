package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))
	//capitalizes each word
	fmt.Println(strings.Title("head first go"))
	fmt.Printf("hello, new line %s\n", "Abhilash")
	fmt.Printf("hello, tab %s\t", "Abhilash\n")
	fmt.Printf("hello, vertical tab %s\v", "Abhilash")
	fmt.Println("Quotes \"\"")
	fmt.Println("A Line", "my guy")
	fmt.Println("Another line\n")
	fmt.Println("Carriage Return\f", "test")
	fmt.Println('A')
	fmt.Printf("hello, decimal %f\n", 2.75)
	fmt.Printf("hello, number %d\n", 2)
	fmt.Printf("hello, bool %t\n", true)
	fmt.Printf("hello, binary %b\n", 256)
	fmt.Printf("hello, numbers %d\n", 50)
	fmt.Printf("hello addition, %d + %d = %d\n", 50, 50, 50+50)
	fmt.Printf("hello subtraction, %d - %d = %d\n", 50, 50, 50-50)
	fmt.Printf("hello division, %d / %d = %d\n", 50, 50, 50/50)
	fmt.Printf("hello multiplication, %d * %d = %d\n", 50, 50, 50*50)
	fmt.Printf("hello equals, %d == %d = %t\n", 50, 50, 50 == 50)
	fmt.Printf("hello does not equals, %d != %d = %t\n", 50, 50, 50 != 50)
	fmt.Printf("hello greater than %d > %d = %t\n", 50, 50, 50 > 50)
	fmt.Printf("hello less than %d < %d = %t\n", 50, 50, 50 < 50)
	fmt.Printf("hello greater than not equals, %d >= %d = %t\n", 50, 50, 50 >= 50)
	fmt.Printf("hello less than or equal, %d <= %d = %t\n", 50, 50, 50 <= 50)
	fmt.Println("hello, Int type: ", reflect.TypeOf(256))
	fmt.Println("hello, float64 type: ", reflect.TypeOf(24.42))
	fmt.Println("hello, Boolean type: ", reflect.TypeOf(true))
	fmt.Println("hello, String type: ", reflect.TypeOf("hi hows't going"))

	// can als just declare without assigning
	var quantity int = 4
	var length, width float64 = 1.2, 2.4
	var customerName = "Abhilash Nair"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Printf("each with an area of %.2f\n", length*width)
	fmt.Println("var Quantity is of type: ", reflect.TypeOf(quantity))
	fmt.Println("var Customer is of type: ", reflect.TypeOf(customerName))
	fmt.Println("var area is of type: ", reflect.TypeOf(length*width))

	var myInt int
	var myFloat float64
	var myString string
	var myBool bool
	fmt.Println("Int's zero value is", myInt)
	fmt.Println("Float's zero value is", myFloat)
	fmt.Println("String's zero value is", myString)
	fmt.Println("Bool's zero value is", myBool)

	//Conversions
	length2 := 3.75
	width2 := 2
	fmt.Println("Area is", length2*float64(width2))
	fmt.Println(float64(width2))
	width2 = int(length2)
	fmt.Println(width2)

}
