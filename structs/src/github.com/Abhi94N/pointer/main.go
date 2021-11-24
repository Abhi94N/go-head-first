package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

type part struct {
	description string
	count       int
}

func doublePack(p *part) {
	p.count *= 2
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func createCar(name string) *car {
	c := car{name: name, topSpeed: 225}
	return &c
}

func main() {

	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)

	nisaan350z := createCar("Nissan 350z")
	fmt.Println(nisaan350z)
	nitroBoost(nisaan350z)
	fmt.Println(*nisaan350z)

	var fuses part
	fuses.description = "Fuses"
	fuses.count = 6
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)

	bolt := &part{description: "Bolt", count: 4}
	fmt.Println(bolt)
	doublePack(bolt)
	fmt.Println(*bolt)
	fmt.Println(&bolt)

}
