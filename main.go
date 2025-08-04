package main

import (
	"fmt"

)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Welcome to Zippy's Calculator!")
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	var result float64

	if operator == "+" {
		result = num1 + num2
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "*" {
		result = num1 * num2	
	} else if operator == "/" {
		result = num1 / num2
	} else {
		fmt.Println("Invalid operator!")
		return
	}

	fmt.Println("Result:", result)
}