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
	grade := promptForGrade()
	grade_status := status(grade)
	fmt.Println(grade)
	fmt.Println("A grade of", grade, "is", grade_status)

}

func promptForGrade() float64 {
	fmt.Print("Enter a grade: ")
	// Set up a buffered reader that gets test from key board
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reflect.TypeOf(reader))
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
	return grade
}

func status(grade float64) string {
	if grade == 100 {
		return "Perfect"
	} else if grade >= 60 {
		return "passing"
	} else {
		return "failing"
	}
}
