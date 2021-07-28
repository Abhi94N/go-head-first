package main

import (
	"fmt"

	"github.com/Abhi94N/error"
	"github.com/Abhi94N/gadget"
	"github.com/Abhi94N/mypkg"
	"github.com/Abhi94N/noisemaker"
	"github.com/Abhi94N/stringer"
	"github.com/Abhi94N/vehicle"
)

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {

	var player gadget.Player = gadget.TapePlayer{}
	mixtape := []string{"Bad Blood", "Oh my good who is she", "Some other Taylor Swift song"}
	playList(player, mixtape)
	gadget.TryOut(player)
	player = gadget.TapeRecorder{}

	playList(player, mixtape)
	gadget.TryOut(player.(gadget.TapeRecorder))

	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}

	s := gadget.Switch("off")
	var t gadget.Toggleable = &s
	t.Toggle()
	fmt.Println(s)
	t.Toggle()
	fmt.Println(s)

	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())

	var toy noisemaker.NoiseMaker
	toy = noisemaker.Whistle("Toyco Canary")
	toy.MakeSound()
	noisemaker.Play(toy)
	toy = noisemaker.Horn("Toyco Blaster")
	toy.MakeSound()
	noisemaker.Play(toy)
	toy = noisemaker.Robot("Optimus Prime")
	toy.MakeSound()
	var robot noisemaker.Robot = toy.(noisemaker.Robot)
	robot.Walk()

	var ride vehicle.Vehicle = vehicle.Car("Toyota Camry")
	ride.Accelerate()
	ride.Steer("left")

	ride = vehicle.Truck("Toyota Tundra")
	ride.Brake()
	ride.Steer("right")

	vehicle.TryVehicle(vehicle.Truck("Ford F150"))

	var err error.Error
	err = error.ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)

	err = error.CheckTemperature(121.379, 100)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	coffeePot := stringer.CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())

	fmt.Println(stringer.Gallons(12.0924832))
	fmt.Println(stringer.Liters(12.0924832))
	fmt.Println(stringer.Milliliters(12.0924832))

	AcceptAnything(3.14159)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(noisemaker.Whistle("Whistle"))

}

func AcceptAnything(thing interface{}) {
	//fmt package has empty interfaces and therefore acceptts anything
	fmt.Println(thing)
	whistle, ok := thing.(noisemaker.Whistle)
	if ok {
		whistle.MakeSound()
	}
}
