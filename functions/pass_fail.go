package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	grade_status := status(grade)
	fmt.Println(grade)
	fmt.Println("A grade of", grade, "is", grade_status)

	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Fahrenheit\n", fahrenheit)
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}

func getFloat() (float64, error) {

	// Set up a buffered reader that gets test from key board
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("\n", reflect.TypeOf(reader))
	// Read until the user pressed enter
	// is multi return values with second being err but if err variable is not used. set _ placeholder
	// input, err := reader.ReadString('\n')
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	//short variable delcleration as long as its declared like the line above will be allowed
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	//Error Handling
	if err != nil {
		return 0, err
	}
	return number, nil
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
