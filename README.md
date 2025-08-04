# Zippy’s Calculator

A beginner-friendly command-line calculator written in Go.  
You type two numbers and pick an operation; it prints the answer—simple as that.

---

## TL;DR

```bash
git clone <repo-url> zippy-calc   # or just copy the three *.go files
cd zippy-calc
go mod init zippy-calc            # once, to create go.mod
go run .                          # launch the calculator
```

---

## Why this project exists

- Learn Go basics – packages, the `main()` function, `go mod`.
- See clean project structure – logic, UI, and `main` split into files.
- Practice reading annotated code – every single line is explained in plain English.

---

## Project layout

```
zippy-calc/
├── main.go          # program entry point – prints banner, starts UI loop
├── ui.go            # text-based menu, user input, validation
└── calculator.go    # pure math functions + error handling
```

All three source files share the same package name (`package main`), so they compile into one executable.

---

## Prerequisites

- **Go 1.22 or newer** – download from <https://golang.org/dl> and install.

  ```bash
  go version   # should show go1.22 or later
  ```

- **A terminal** (Command Prompt, PowerShell, Bash, Zsh…) – built-in on macOS, Linux, and Windows.

---

## Getting started

1. Clone or copy the `zippy-calc` folder to any directory.  
2. Initialize a module (creates `go.mod`):

   ```bash
   go mod init zippy-calc
   ```

3. Run the app:

   ```bash
   go run .
   ```

### Example session

```
=== Zippy's Calculator v2 ===
Choose an operation:
  1) +
  2) -
  3) *
  4) /
  5) Exit
Select option:
```

4. Build an executable (optional):

   ```bash
   go build -o zippy
   ./zippy
   ```

---

## How the code is organised

| File            | Responsibility                                                   | Beginner take-aways                                   |
| --------------- | ---------------------------------------------------------------- | ----------------------------------------------------- |
| `main.go`       | Starts the program, prints a banner, calls the UI loop.          | Shows the minimal structure of a Go `main` package.   |
| `ui.go`         | Handles all user interaction (menu, number input, validation).   | Demonstrates buffered input and string→number parsing |
| `calculator.go` | Pure math functions plus error checking (division by zero).      | Separates business logic from I/O; shows error returns|

---

## Extending the calculator

- **Add new operators**

  1. Append the symbol to `ops` in `ui.go`.  
  2. Add a `case` in `Calculate()` inside `calculator.go` (or write a helper).

- **Make it GUI-based** — swap the console UI for a toolkit like [Fyne](https://fyne.io) or [Gio](https://gioui.org).

- **Add unit tests** — create `calculator_test.go` and test each math helper with Go’s `testing` package.

---

## Contributing

Pull requests are welcome—especially if you:

- Spot a typo in the comments
- Improve input validation
- Add more beginner-friendly notes

---

## License

MIT – do whatever you want, just keep the copyright notice.

---

Happy hacking! ✨