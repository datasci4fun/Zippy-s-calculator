// Same package so we can share functions and variables easily.
package main

// Packages that help us read input and convert strings to numbers.
import (
	"bufio"   // Reads user input line by line.
	"fmt"     // Prints text.
	"os"      // Gives us access to the keyboard and screen.
	"strconv" // Turns text like "42" into the number 42.
	"strings" // Helps trim spaces and new-lines.
)

// A list of math symbols we allow.
var ops = []string{"+", "-", "*", "/"}

// This function keeps showing a menu until the user chooses “Exit”.
func StartInteractiveLoop() {
	// Create one reader we’ll reuse for every question.
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop—breaks only when user exits.
	for {
		// 1) Show the menu.
		fmt.Println()
		fmt.Println("Choose an operation:")
		for i, op := range ops {
			fmt.Printf("  %d) %s\n", i+1, op) // e.g. “ 1) + ”
		}
		// Extra choice: Exit.
		fmt.Printf("  %d) Exit\n", len(ops)+1)

		// 2) Read the user’s menu choice.
		fmt.Print("Select option: ")
		choiceText, _ := reader.ReadString('\n') // read until Enter key
		choiceText = strings.TrimSpace(choiceText)

		// Turn the text (like "2") into a real number.
		choice, err := strconv.Atoi(choiceText)
		// If the text wasn’t a number or is out-of-range, start over.
		if err != nil || choice < 1 || choice > len(ops)+1 {
			fmt.Println("⚠️  Invalid menu option.")
			continue
		}
		// If they picked “Exit”, leave the loop and end program.
		if choice == len(ops)+1 {
			fmt.Println("Good-bye!")
			return
		}

		// Get the matching operator symbol.
		operator := ops[choice-1]

		// 3) Ask for the first number.
		num1, ok := readFloat(reader, "Enter first number: ")
		if !ok { // they typed something that isn’t a number
			continue
		}

		// 4) Ask for the second number.
		num2, ok := readFloat(reader, "Enter second number: ")
		if !ok {
			continue
		}

		// 5) Do the math (function lives in calculator.go).
		result, err := Calculate(num1, num2, operator)
		if err != nil { // e.g. division by zero
			fmt.Println("⚠️ ", err)
			continue
		}

		// 6) Show the answer.
		fmt.Printf("Result: %.6g\n", result)
	}
}

// Helper: ask a question, read text, turn it into a float64.
// Returns (value, true) on success, (_, false) if the user messed up.
func readFloat(r *bufio.Reader, prompt string) (float64, bool) {
	fmt.Print(prompt)
	text, _ := r.ReadString('\n')
	text = strings.TrimSpace(text)

	// Convert the text to a floating-point number.
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("⚠️  Not a valid number.")
		return 0, false
	}
	return val, true
}