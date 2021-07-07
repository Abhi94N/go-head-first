package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a grade: ")
	// Set up a buffered reader that gets test from key board
	reader := bufio.NewReader(os.Stdin)
	// Read until the user pressed enter
	// is multi return values with second being err but if err variable is not used. set _ placeholder
	// input, err := reader.ReadString('\n')
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	//short variable delcleration as long as its declared like the line above will be allowed
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if int(grade) == 100 {
		status = "Perfect"
	} else if int(grade) >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(grade)
	fmt.Println("A grade of", grade, "is", status)
	fmt.Println(reflect.TypeOf(reader))

}
