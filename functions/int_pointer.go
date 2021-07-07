package main

import "fmt"

func main() {

	myInt := 7
	myIntPointer := &myInt

	fmt.Println(myInt)
	fmt.Println(*myIntPointer)
	fmt.Println(&myInt)
	fmt.Println(myIntPointer)

}
