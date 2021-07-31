package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("sum: %0.2f\n", sum)

	err = Socialize()
	if err != nil {
		log.Fatal(err)
	}

}

func Socialize() error {
	// will run last
	defer fmt.Println("Goodbye")
	fmt.Println("Hello!")

	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Noice weather eh?")
	return nil
}

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing File", file.Name())
	file.Close()
}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64

	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	//defer this call to be called at the end of the function close to handle any error closing
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)

	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
