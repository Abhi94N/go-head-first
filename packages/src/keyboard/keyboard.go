// Package keyboard reads user input from the keyboard
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered
func GetFloat() (float64, error) {

	// Set up a buffered reader that gets test from key board
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("\n", reflect.TypeOf(reader))
	// Read until the user pressed enter
	// is multi return values with second being err but if err variable is not used. set _ placeholder
	// input, err := reader.ReadString('\n')
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	//short variable delcleration as long as its declared like the line above will be allowed
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	//Error Handling
	if err != nil {
		return 0, err
	}
	return number, nil
}
