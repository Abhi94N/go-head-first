package main

import "fmt"

func main() {

	//instead of int
	var count int = 12
	// instead of append
	var suffix string = "minutes of bonus footage"
	//instead of fmt
	var format string = "DVD"
	var languages = append([]string{}, "Espanol", "English")
	fmt.Println(count, suffix, "on", format, languages)
}
