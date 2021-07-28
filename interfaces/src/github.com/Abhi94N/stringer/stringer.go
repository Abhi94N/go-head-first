package stringer

import "fmt"

type Stringer interface {
	String() string
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

//you can [ass any stringer function to fmt.print, println, printf

//set type to set up interface with
type Gallons float64

//satisfy the interface
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

//set type to set up interface with
type Liters float64

//satisfy the interface
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

//set type to set up interface with
type Milliliters float64

//satisfy the interface
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}
