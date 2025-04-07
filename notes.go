package main

import "fmt"

func main() {
	var notes []string
	entry := "NOTES"
	for entry != "EXIT" {
		fmt.Print("Write your entry: ")
		fmt.Scanln(&entry)
		if entry != "EXIT" {
			notes = append(notes, entry)
		}
	}
	fmt.Println("Your Notes: ", notes)
}
