package main

import (
	"fmt"
	"math"
)

func main() {

	severalInts(1)

	severalStrings("a", "b")

	fmt.Println(maximum(15, 51, 231, 1234.0))

	fmt.Println(inRange(7, 23, -12, 20, 8, 9, 10, 6))

	fmt.Println(average(5, 5, 5, 5, 10))
	fmt.Println(sum(5, 5, 5, 5, 10))

}

func sum(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func average(numbers ...float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func maximum(numbers ...float64) float64 {
	//set max to negative infinity
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max

}

func severalInts(numbers ...int) {
	fmt.Println(numbers)

}
func severalStrings(strings ...string) {
	fmt.Println(strings)

}
