package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	// time package has Time type that represents the time
	var now time.Time = time.Now()
	year := now.Year()
	fmt.Println(year)

	//Replacer type to replace value
	broken := "G# r#ckets!"
	replacer := strings.NewReplacer("#", "o")
	var replacer2 *strings.Replacer = strings.NewReplacer("!", "?")
	fixed := replacer.Replace(broken)
	fmt.Println(broken, fixed)
	fmt.Println(reflect.TypeOf(replacer))
	fmt.Println(replacer2.Replace(fixed))

	//Loops
	x := 1
	for x <= 3 {
		fmt.Println(x)
		x++
	}

	var y int
	for y = 3; y > 0; {
		fmt.Println(y)
		y--
	}

}
