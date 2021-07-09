package main

import (
	"fmt"
	// "greeting"
	// "greeting/dansk"
	// "greeting/deutsch"
	"log"

	"github.com/Abhi94N/keyboard"
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/dansk"
	"github.com/headfirstgo/greeting/deutsch"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	grade_status := status(grade)
	fmt.Println(grade)
	fmt.Println("A grade of", grade, "is", grade_status)

	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Fahrenheit\n", fahrenheit)
	fmt.Printf("%0.2f degrees Celsius\n", celsius)

	greeting.Hello()
	dansk.Hej()
	deutsch.Hallo()

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
