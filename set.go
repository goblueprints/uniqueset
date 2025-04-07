package uniqueset

// Set - collection of unique values of any comparable type
type Set[T comparable] struct {
	data map[T]struct{}
}

// New creates a new Set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

// Add adds a value to the set
func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

// AddAll adds multiple values to the set
func (s *Set[T]) AddAll(values ...T) {
	for _, v := range values {
		s.Add(v)
	}
}

// Remove deletes a value from the set
func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

// Contains checks if a value exists in the set
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// Values returns all values as a slice
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for v := range s.data {
		values = append(values, v)
	}
	return values
}

// Size returns the number of elements in the set
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the set
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Union combines two sets
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	for v := range s.data {
		result.Add(v)
	}
	for v := range other.data {
		result.Add(v)
	}
	return result
}

// Intersection returns common elements between two sets
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	for v := range s.data {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns elements in s that are not in other (s - other)
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for v := range s.data {
		if !other.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

// Equal checks if two sets contain exactly the same elements
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for v := range s.data {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}
