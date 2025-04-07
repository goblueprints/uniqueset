package main

import (
	"fmt"

	"github.com/goblueprints/uniqueset"
)

type Point struct {
	X, Y int
}

func main() {
	pointSet := uniqueset.New[Point]()
	pointSet.Add(Point{1, 2})
	pointSet.Add(Point{3, 4})

	fmt.Println(pointSet.Contains(Point{1, 2})) // true
}
