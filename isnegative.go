package main

import "github.com/01-edu/z01"

// IsNegative checks if n is negative and prints "N" or "P"
func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

// main is the entry point when you do `go run`
func main() {
	IsNegative(-1) // try with -1
	IsNegative(0)  // try with 0
	IsNegative(5)  // try with 5
}
