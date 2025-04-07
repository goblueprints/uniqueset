package main

import (
	"fmt"

	"github.com/goblueprints/uniqueset"
)

func main() {
	// Create a new set
	set := uniqueset.New[int]()

	// Add elements
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(2) // Duplicate - won't be added

	// Check size
	fmt.Println("Size:", set.Size()) // 3

	// Check membership
	fmt.Println("Contains 2:", set.Contains(2)) // true

	// Get all values
	fmt.Println("Values:", set.Values()) // [1 2 3] (order may vary)

	// Remove element
	set.Remove(2)
	fmt.Println("After removal:", set.Values()) // [1 3]
}
