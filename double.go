package main

import "fmt"

func doubleNum(number int) int {
	return (number * 2)
}

func main() {
	num := 10
	fmt.Printf("Original number:%d\n", num)
	fmt.Printf("Doubled number:%d\n", doubleNum(num))
}
