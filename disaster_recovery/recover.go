package main

import "fmt"

func calmDown() {
	p := recover()
	//returns error due to returning an empty interface which leads to compile
	//fmt.Println(p.Error())
	//
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func freakOut() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
	recover()
	fmt.Println("I will never run")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
