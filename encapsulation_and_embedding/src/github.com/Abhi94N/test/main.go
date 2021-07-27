package main

import (
	"fmt"
	"log"

	"github.com/Abhi94N/calendar"
	"github.com/Abhi94N/geo"
)

func main() {

	date := calendar.Date{}
	fmt.Println(date)
	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year())
	err = date.SetMonth(7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Month())
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Day())

	coordinates := geo.Coordinates{}

	err = coordinates.SetLatitude(89)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	err = coordinates.SetLongitude(137)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Longitude())

	event := calendar.Event{}
	fmt.Println(event)
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())

	err = event.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Month())

	err = event.SetDay(11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Day())
	err = event.SetTitle("Drake & Josh")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())

	location := geo.Landmark{}
	err = location.SetName("King and Queen Tower")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name())

	err = location.SetLatitude(84.355)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Latitude())

	err = location.SetLongitude(33.916)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Longitude())
}
