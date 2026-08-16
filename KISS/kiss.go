package main

import "fmt"

// Calculator defines a simple calculator structure
type Calculator struct{}

// Add method adds two numbers
func (c Calculator) Add(a, b int) int {
	return a + b
}

// Subtract method subtracts two numbers
func (c Calculator) Subtract(a, b int) int {
	return a - b
}

func main() {
	calculator := Calculator{}

	// Calculate 5 + 3
	result1 := calculator.Add(5, 3)
	fmt.Println("5 + 3 =", result1)

	// Calculate 8 - 2
	result2 := calculator.Subtract(8, 2)
	fmt.Println("8 - 2 =", result2)
}
