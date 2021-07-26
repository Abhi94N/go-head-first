package main

import (
	"fmt"

	"github.com/Abhi94N/geo"
	"github.com/Abhi94N/magazine"
)

type student struct {
	name  string
	grade float64
}

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

type Liters float64
type Mililiters float64
type Gallons float64

type Title string

type Population int

type MyType string

type Number int

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
	fmt.Printf("%#v\n", myStruct)

	var subscriber1 magazine.Subscriber

	subscriber1.Name = "Abhilash Nair"
	subscriber1.Rate = 500.0
	subscriber1.Active = true
	fmt.Println("Name:", subscriber1.Name)
	fmt.Println("Name:", subscriber1.Rate)
	fmt.Println("Name:", subscriber1.Active)

	var subscriber2 magazine.Subscriber
	subscriber2.Name = "Aman Singh"
	fmt.Println(subscriber2)

	subscriber3 := magazine.Subscriber{Name: "Test"}
	fmt.Println(subscriber3)

	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top Speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
	bolts.count = 15
	showInfo(bolts)

	nails := minimumOrder("Nails")
	showInfo(nails)

	//now a subscriber pointer
	abhi := defaultSubscriber("Abhiash Nair")
	printInfo(abhi)
	applyDiscounts(abhi)
	printInfo(abhi)

	sarah := newStudent("Sarah Shahi")
	printStudent(sarah)

	var value int = 2
	var pointer *int = &value
	fmt.Println(*pointer)

	//pointer referencing value
	var studentPointer *student = &sarah
	fmt.Println((*studentPointer).name)
	studentPointer.grade = 97
	fmt.Println(studentPointer)
	fmt.Println(*studentPointer)

	address := magazine.Address{Street: "789 Hammond Drive", City: "Atlanta", State: "GA", PostalCode: "30062"}
	fmt.Printf("%s %s %s, %s\n", address.Street, address.City, address.State, address.PostalCode)

	//location
	location := geo.Coordinates{Longitude: -122.08, Latitude: 37.08}
	fmt.Println("Latitude", location.Latitude)
	fmt.Println("Longitude", location.Longitude)

	employee := magazine.Employee{Name: "Abhilash Nair", Salary: 180000, Address: address}
	fmt.Println(employee.Name)
	fmt.Println("Salary", employee.Salary)
	fmt.Println("Addresss", employee.Address)
	fmt.Println("Postal Code", employee.PostalCode)

	subscriber1.Address = address

	fmt.Println(employee)
	fmt.Println(subscriber1)
	fmt.Printf("%#v\n", subscriber2.Address)
	fmt.Printf("%#v\n", subscriber1.City)

	//
	location2 := geo.Landmark{Name: "King and Queen Tower"}
	location2.Longitude = 33.916
	location2.Latitude = 84.355
	fmt.Println(location2)

	carFuel := Liters(240.0)
	busFuel := Gallons(10.0)
	fmt.Println(carFuel, busFuel)

	//to convert one type to another
	carFuel_2 := Gallons(carFuel * 0.264)
	busFuel_2 := Liters(busFuel * 3.785)

	GallonsInMililters := GallonsToMililiters(carFuel_2)
	MililitersInGallons := MililitersToGallons((GallonsInMililters))

	LitersInMililiters := LitersToMililiters(busFuel_2)
	MililitersInLiters := MililitersToLiters(LitersInMililiters)

	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel_2, busFuel_2)
	fmt.Println("Gallons in mililiters: ", GallonsInMililters, "mililiters in gallons:", MililitersInGallons)
	fmt.Println("liters in mililiters: ", LitersInMililiters, "mililiters in liters:", MililitersInLiters)

	//converting types using functions
	carFuel_2 += LitersToGallons(Liters(40.0))
	busFuel_2 += GallonsToLiters(Gallons(30.0))

	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel_2)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel_2)

	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	//asci order
	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Zodiac"))
	fmt.Println(Title("Alien") > Title("Zodiac"))
	fmt.Println(Title("Alien") + "s")
	//below will fail because it is an invalid operator
	// fmt.Println(Title("Jaws 2") - " 2")
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Liters(1.2) == 1.2)
	fmt.Println(Liters(1.2) < 3.4)

	population := Population(572)
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congratulations, Levin and Anna! IT's a girl")
	population += 1
	fmt.Println("Sleepy Creek County population:", population)

	values := MyType("A MyType value")
	values.sayHi()
	anothervalue := MyType("Another value")
	anothervalue.sayHi()

	values_2 := MyType("My Type 2 Value")
	values_2.MethodWithParameters(30, true)

	fmt.Println(values_2.WithReturn(" Returning"))

	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	fmt.Println("Ten before doubling:", ten)
	ten.Double()
	fmt.Println("Ten after doubling:", ten)
	fmt.Println("Four before doubling:", four)
	four.Double()
	fmt.Println("Four after doubling:", four)

	value_3 := MyType("a value")
	pointer_3 := &value_3
	value_3.method()
	value_3.pointerMethod()
	pointer_3.method()
	pointer_3.pointerMethod()

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals to %0.3f gallons\n", soda, soda.ToGallons())
	fmt.Printf("%0.3f liters equals to %0.3f mililiters\n", soda, soda.ToMililiters())
	water := Mililiters(500)
	fmt.Printf("%0.3f Mililiters equals to %0.3f gallons\n", water, water.ToGallons())
	fmt.Printf("%0.3f Mililiters equals to %0.3f liters\n", water, water.ToLiters())
	gas := Gallons(20)
	fmt.Printf("%0.3f gallons equals to %0.3f Mililiters\n", gas, gas.ToMililiters())
	fmt.Printf("%0.3f gallons equals to %0.3f Liters\n", gas, gas.ToLiters())

}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)

}

func (m Mililiters) ToLiters() Liters {
	return Liters(m / 1000)

}

func (g Gallons) ToMililiters() Mililiters {
	return Mililiters(g * 3785.41)

}

func (l Liters) ToMililiters() Mililiters {
	return Mililiters(l * 1000)

}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Mililiters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func (n *Number) Double() {
	*n *= 2
}

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

// method with parameters
func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

//methods with retirh value

func (m MyType) WithReturn(s string) MyType {
	return m + MyType(s)
}

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

// function withd efined types
func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)

}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * .264)

}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)

}

func MililitersToGallons(m Mililiters) Gallons {
	return Gallons(m * .000264)

}

func GallonsToMililiters(g Gallons) Mililiters {
	return Mililiters(g * 3785.41)

}

func MililitersToLiters(m Mililiters) Liters {
	return Liters(m * .001)

}

func LitersToMililiters(l Liters) Mililiters {
	return Mililiters(l * 1000)

}

//update to take a pointer
func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly Rate:", s.Rate)
	fmt.Println("Active?:", s.Active)
}

func printStudent(s student) {
	fmt.Println("Name:", s.name)
	fmt.Println("Grade:", s.grade)
}

func newStudent(name string) student {
	s := student{name: name, grade: 87.0}
	return s
}

//update to return a pointer
func defaultSubscriber(name string) *magazine.Subscriber {
	s := magazine.Subscriber{Name: name, Rate: 5.99, Active: true}
	// REturns a pointer to a struct isntead of the struct itself
	return &s
}

//modifying a struct using a function. must pass address to copy original
func applyDiscounts(s *magazine.Subscriber) {
	s.Rate = 4.99
}
