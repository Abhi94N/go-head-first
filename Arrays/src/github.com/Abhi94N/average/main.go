package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	var notes [7]string

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])
	fmt.Println(notes)

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[1])

	var dates [3]time.Time
	dates[0] = time.Unix(time.Now().Unix(), 0)
	dates[1] = time.Unix(1508632200, 0)
	dates[2] = time.Unix(1447920000, 0)
	fmt.Println(dates[0])
	fmt.Println(dates[1])
	fmt.Println(dates[2])

	//zero values
	var primes_2 [5]int
	primes_2[0] = 0

	fmt.Println(primes[0])
	fmt.Println(primes[2])
	fmt.Println(primes[4])

	var notes_2 [7]string
	notes_2[0] = "do"
	fmt.Println(notes[6])
	fmt.Println(notes[3])
	fmt.Println(notes[0])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	var notes_3 [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes_3[3], notes_3[6], notes_3[0])
	var primes_3 [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes_3[0], primes_3[2], primes_3[4])

	//short version
	notes_4 := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes_4 := [5]int{2, 3, 5, 7, 11}
	fmt.Println(notes_4[2], notes_4[4], notes_4[5])
	fmt.Println(primes_4[1], primes_4[3], primes_3[0])
	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}
	fmt.Println(text)

	//shows Array literals
	fmt.Printf("%#v\n", notes_4)
	fmt.Printf("%#v\n", primes_4)

	index := 1
	fmt.Println(index, notes_4[index])
	index = 3
	fmt.Println(index, notes_4[index])

	//loop
	for i := 0; i < len(notes_4); i++ {
		fmt.Println(i, notes_4[i])
	}

	for index, _ := range notes_4 {
		fmt.Println(index)
	}
	for _, notes_4 := range notes_4 {
		fmt.Println(notes_4)
	}

	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	numbers_2 := [6]int{3, 16, -2, 10, 23, 12}
	for i, number_2 := range numbers_2 {
		if number_2 >= 10 && number_2 <= 20 {
			fmt.Println(i, number_2)
		}
	}

	numbers_3, err := datafile.GetFloats("data.txt")
	fmt.Println(numbers_3)
	if err != nil {
		log.Fatal(err)
	}
	var sum_3 float64
	for _, number_3 := range numbers_3 {
		fmt.Println(number_3)
		sum_3 += number_3
	}
	fmt.Printf("Average From file: %0.2f\n", sum_3/float64(len(numbers_3)))

}
