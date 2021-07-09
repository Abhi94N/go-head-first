package main

import (
	"dates"
	"fmt"
	"greeting"
	"greeting/dansk"
	"greeting/deutsch"
	"keyboard"
	"log"
	"math"
)

func main() {

	days := 45
	weeks := int(math.Floor(dates.DaysToWeeks(days)))
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
	fmt.Println("That is after", weeks, "weeks and", (days % weeks), "days")

	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
	dansk.Hej()
	dansk.GodTag()

	greeting.AllGreetings()

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
