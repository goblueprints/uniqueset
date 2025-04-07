package main

import (
	"fmt"

	"github.com/goblueprints/uniqueset"
)

func main() {
	setA := uniqueset.New[string]()
	setA.AddAll("apple", "banana", "orange")

	setB := uniqueset.New[string]()
	setB.AddAll("banana", "kiwi", "pear")

	// Union
	union := setA.Union(setB)
	fmt.Println("Union:", union.Values()) // ["apple", "banana", "orange", "kiwi", "pear"]

	// Intersection
	intersection := setA.Intersection(setB)
	fmt.Println("Intersection:", intersection.Values()) // ["banana"]

	// Difference
	difference := setA.Difference(setB)
	fmt.Println("Difference:", difference.Values()) // ["apple", "orange"]
}
