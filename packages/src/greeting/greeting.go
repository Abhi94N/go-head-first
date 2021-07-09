package greeting

import (
	"fmt"
)

//Capitalized functions can be exported
func Hello() {
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi!")
}

//N0 Package qualifier needed
func AllGreetings() {
	Hello()
	Hi()
}
