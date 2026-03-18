// High-Performance Cryptographic Hashing Tool
// Go has a special feature called Go Assembly (Plan 9). It allows us to write functions in .s files and call them directly from .go files.
// This is how the Go standard library makes math and crypto so fast . . . SO this is very simple project (as an entry point) to implement it
package main

import "fmt"

func ProcessBlock(data []byte)

func main() {
	input := []byte("HELLO ASSEMBLY!")
	fmt.Printf("Before %s\n", input)
	ProcessBlock(input)
	fmt.Printf("After %s\n", input)
}