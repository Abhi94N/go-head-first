package main

import (
	"fmt"
)

func main() {

	var ranks map[string]int

	ranks = make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3

	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}

	var ok bool
	var rank int
	//can allow negative values
	rank, ok = ranks["bronze"]

	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	fmt.Println(ranks)
	fmt.Println(ranks["silver"])

	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"

	fmt.Println(elements["Li"])
	fmt.Println(elements)

	isPrime := make(map[int]bool)

	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])
	isPrime[5] = true
	var prime bool
	prime, ok = isPrime[7]
	fmt.Printf("prime: %v, ok %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok %v\n", prime, ok)

	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap)

	ranks_2 := map[string]int{"gold": 1, "silver": 2, "bronze": 3}
	fmt.Println(ranks_2)

	elements_2 := map[string]string{"H": "Hydrogen", "Li": "Lithium"}
	fmt.Println(elements_2)

	// empty map
	emptyMap := map[string]float64{}
	fmt.Println(emptyMap)

	//Zero values
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 7
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I've not been assigned"])

	words := make(map[string]string)
	words["I've been assigned"] = "hello"
	fmt.Printf("%#v\n", words["I've been assigned"])
	fmt.Printf("%#v\n", words["I've not been assigned"])

	counters := make(map[string]int)

	counters["a"]++
	counters["a"]++
	counters["b"]++
	counters["c"]++

	fmt.Println(counters)
	fmt.Println(counters["a"], counters["b"], counters["c"])

	var value int
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	value, ok = counters["d"]
	fmt.Println(value, ok)

	// map is nil so you can't add a value because it is not made or assigned
	// var nilMap map[int]string
	// fmt.Printf("%#v\n", nilMap)
	// nilMap[3] = "three"

	myMaps := make(map[int]string)
	myMaps[3] = "three"
	fmt.Printf("%#v\n", myMaps)

	//
	status("Jesus")
	status("Buddy Guy")
	status("Al Davis")

	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)

	for _, item := range data {
		counts[item]++
	}

	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}

	// Map for range loop
	grades := map[string]float64{"Al Davis": 0, "Abhilash Nair": 100, "Buddy Guy": 40, "Jo Francis": 40.2}
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}

	var names []string

	//key
	fmt.Println("Class Roster")
	for name := range grades {
		fmt.Println(name)
		names = append(names, name)
	}

	fmt.Println("Grades")
	for _, grade := range grades {
		fmt.Println(grade)
	}

	//go through names list
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}

}

func status(name string) {
	grades := map[string]float64{"Al Davis": 0, "Abhilash Nair": 100, "Buddy Guy": 40, "Jo Francis": 40.2}
	grade, ok := grades[name]

	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing with a %#v \n", name, grade)
	}
}
