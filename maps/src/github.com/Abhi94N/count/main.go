package main

import (
	"fmt"
	"log"

	"github.com/Abhi94N/datafile"
)

func main() {

	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
	var counts []int
	var names []string

	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true

			}

		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}

	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}

	results := countVotes(lines)
	fmt.Println(results)

	for name, count := range results {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}

func countVotes(votes []string) map[string]int {

	counts := make(map[string]int)
	for _, vote := range votes {
		counts[vote]++
	}

	return counts
}
