package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Charlie": 95,
	}

	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}
}