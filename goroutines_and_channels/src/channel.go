package main

import (
	"fmt"
	"time"
)

//Take chan as a parameter
func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")

}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)

	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()

	anotherChannel := make(chan string)
	go send(anotherChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-anotherChannel)
	fmt.Println(<-anotherChannel)
}
