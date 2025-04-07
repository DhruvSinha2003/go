package main

import "fmt"

func calculate(a, b float64, op string) {
	switch op {
	case "+":
		result := a + b
		fmt.Print(result)

	case "-":
		result := a - b
		fmt.Print(result)

	case "*":
		result := a * b
		fmt.Print(result)

	case "/":
		result := a / b
		fmt.Print(result)
	}
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	calculate(num1, num2, operator)
}
