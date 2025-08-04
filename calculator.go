// Same package name so these functions are visible to ui.go and main.go.
package main

import "errors" // Used to create error messages.

// Calculate decides which math operation to run.
func Calculate(a, b float64, op string) (float64, error) {
	switch op { // simpler than a long if-else chain
	case "+":
		return add(a, b), nil
	case "-":
		return sub(a, b), nil
	case "*":
		return mul(a, b), nil
	case "/":
		return div(a, b)
	default:
		return 0, errors.New("unknown operator")
	}
}

// The four basic math helpers.
func add(x, y float64) float64 { return x + y }
func sub(x, y float64) float64 { return x - y }
func mul(x, y float64) float64 { return x * y }

// div checks for division by zero and reports an error.
func div(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}