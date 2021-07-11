package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Provide a file: ")
	// Set up a buffered reader that gets test from key board
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	file_path := strings.TrimSpace(input)
	// open file for reading
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	//use scanner to read file
	scanner := bufio.NewScanner(file)

	//loops through scanner and while true will loop
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
