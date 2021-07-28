package noisemaker

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

//Interface requires each type to have a function called make soundd
type NoiseMaker interface {
	MakeSound()
}

//method created for interface
func Play(n NoiseMaker) {
	n.MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Charging up")
}
