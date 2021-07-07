package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//Package variable
var metersPerLiter float64

func main() {

	err := errors.New("height can't be negative")
	fmt.Println(err.Error())
	err = fmt.Errorf("a height of %0.2f is invalid", 2.342)

	fmt.Println(err)
	metersPerLiter = 10.0
	sayHi()
	repeatLine("Im in main", 20)
	var amount, total float64

	amount, err = paintNeeded(20, 30)
	fmt.Printf("%.2f liters\n", amount)
	total += amount

	amount, err = paintNeeded(20, 30.5)
	fmt.Printf("%.2f liters\n", amount)
	total += amount

	amount, err = paintNeeded(5.323, 6.4)
	fmt.Printf("%.2f liters\n", amount)
	total += amount

	fmt.Printf("Total: %0.2f liters\n", total)
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(123123.0))

	// amount, err = paintNeeded(5, -24)
	amount, err = paintNeeded(5, 24)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Printf("%.2f paint liters\n", amount)
	}
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)

	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)

	squareRootValue, err := squareRoot(9.2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Square Root:", squareRootValue)
	}

	quotient, err := divide(9.2, 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}

}

func sayHi() {
	fmt.Println("Say Hi")
	repeatLine("Im in another func called sayHI()", 5)

}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func paintNeeded(width float64, height float64) (float64, error) {
	area := width * height

	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	return area / metersPerLiter, nil
}

func double(number float64) float64 {
	return number * 2
}

func manyReturns() (int, bool, string) {
	return 1, true, "string"
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("Can't get square root of negative number.")
	}
	return math.Sqrt(number), nil
}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}
