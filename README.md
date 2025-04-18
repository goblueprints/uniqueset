# UniqueSet - Generic Set Implementation in Go

A _thread-unsafe_, generic Set implementation for Go 1.18+ that works with any comparable type.

## Features

- Type-safe generic implementation
- All basic set operations:
  - Add/Remove elements
  - Check membership
  - Union, Intersection, Difference
  - Set equality comparison
- Zero dependencies
- Simple API

## Installation

```bash
go get github.com/goblueprints/uniqueset
```

## Usage

### Basic Operations

```go
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
```

### Set Operations

```go
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
```

### Working with Custom Types

```go
type Point struct {
 X, Y int
}

func main() {
 pointSet := uniqueset.New[Point]()
 pointSet.Add(Point{1, 2})
 pointSet.Add(Point{3, 4})
 
 fmt.Println(pointSet.Contains(Point{1, 2})) // true
}

```

## Limitations

1. Not thread-safe (for concurrent use, add synchronization)
2. Values are returned in random order (use sort package if ordering is needed)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
