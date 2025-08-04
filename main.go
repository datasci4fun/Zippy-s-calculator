// Every Go file starts by naming its package.
// The special name “main” means: “this is a runnable program.”
package main

// We only need fmt here. It lets us print text to the screen.
import "fmt"

// main() is where the program begins.
func main() {
	// A friendly hello message.
	fmt.Println("=== Zippy's Calculator v2 ===")

	// Jump into our interactive menu (defined in ui.go).
	StartInteractiveLoop()
}