package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Abhi94N/slice"
)

func main() {

	array := [5]string{"a", "b", "c", "d", "e"}
	sliced := array[1:3]
	sliced = append(sliced, "x")
	sliced = append(sliced, "y", "z")
	for _, letter := range sliced {
		fmt.Println(letter)
	}

	//using slice function in GetFloats
	numbers, err := slice.GetFloats("data.txt")
	fmt.Println(numbers)
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))

	fmt.Printf("%.2f / %.2f = %.2f\n", sum, sampleCount, sum/sampleCount)

	//slices
	var notes []string
	notes = make([]string, 7)

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes)
	fmt.Println(notes[1])

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	fmt.Println(len(notes))
	fmt.Println(len(primes))

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}

	for _, letter := range letters {
		fmt.Println(letter)
	}

	//underlying Array
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("Underlying Array:", underlyingArray)
	fmt.Println(reflect.TypeOf(underlyingArray))
	slice1 := underlyingArray[0:3]
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println("Slice 1", slice1)
	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(reflect.TypeOf(slice2))
	fmt.Println("Slice 2", slice2)
	underlyingArray[4] = "z"

	fmt.Println(underlyingArray[i:])
	fmt.Println(underlyingArray[:j])
	fmt.Println(underlyingArray[:])

	//append
	slice3 := []string{"a", "b"}
	fmt.Println(slice3, len(slice3))
	slice3 = append(slice3, "c", "d")
	fmt.Println(slice3, len(slice3))

	// Creating a larger slice
	s1 := []string{"s1", "s2"}
	s1[0] = "YY"
	s2 := append(s1, "s2", "s2")

	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[1] = "XX"
	fmt.Println(s1, s2, s3, s4)

	fmt.Println(s1, s2, s3, s4)

	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice, boolSlice)

	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)

	fmt.Println("slice length", len(intSlice))

	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v\n", intSlice)

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("slice: %#v\n", slice)

}
