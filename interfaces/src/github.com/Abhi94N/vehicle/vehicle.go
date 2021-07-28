package vehicle

import "fmt"

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding Up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding Up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

//function
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

//assertion to get non interface function to work
func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("Camping Equipment")
	}
}
