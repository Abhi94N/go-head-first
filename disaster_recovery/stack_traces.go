package main

import "fmt"

func one() {
	defer fmt.Println("Deferred in one()")
	two()
}

func two() {
	defer fmt.Println("Deferred in two()")
	three()
}

func three() {
	panic("This call stack's too deep for me")
}

func main() {
	one()
}
