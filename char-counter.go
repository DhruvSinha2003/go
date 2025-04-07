package main

import "fmt"

func main() {
	sentence := "the quick brown fox jumps over the lazy dog"
	var letter string
	fmt.Print("Which letter do you want to count? ")
	fmt.Scan(&letter)

	count := 0

	for _, char := range sentence {
		if string(char) == letter {
			count++
		}
	}
	fmt.Printf("Total count: %d\n", count)
}
