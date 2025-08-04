// ↓ Every Go file starts by declaring its package.
//   “main” marks this file as the program’s entry-point module.
package main

// ↓ Import list. We only need the “fmt” package for
//   formatted I/O (printing to– and reading from– the console).
import (
	"fmt"
)

// ↓ main() is where a standalone Go program begins execution.
func main() {
	// Declare two floating-point operands.
	var num1, num2 float64
	// Declare a string to hold the operator symbol.
	var operator string

	// Friendly banner.
	fmt.Println("Welcome to Zippy's Calculator!")

	// Prompt for the first number…
	fmt.Print("Enter first number: ")
	// …and read it from STDIN.  Scanln() stops at whitespace.
	fmt.Scanln(&num1)

	// Prompt for the operator (+, –, *, /) …
	fmt.Print("Enter operator (+, -, *, /): ")
	// …and read it into the *operator* variable.
	fmt.Scanln(&operator)

	// Prompt for the second number…
	fmt.Print("Enter second number: ")
	// …and read it.
	fmt.Scanln(&num2)

	// Placeholder for the final answer.
	var result float64

	// Branch on the operator and perform the chosen arithmetic.
	if operator == "+" {
		result = num1 + num2
	} else if operator == "-" {
		result = num1 - num2
	} else if operator == "*" {
		result = num1 * num2
	} else if operator == "/" {
		result = num1 / num2
	} else { // Anything else is invalid.
		fmt.Println("Invalid operator!")
		return // Abort early.
	}

	// Display the computed result.
	fmt.Println("Result:", result)
}