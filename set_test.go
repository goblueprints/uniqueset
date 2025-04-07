package uniqueset_test

import (
	"testing"

	"github.com/goblueprints/uniqueset"
)

func TestSet_Int(t *testing.T) {
	s := uniqueset.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(2) // Duplicate!

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	if !s.Contains(1) {
		t.Error("Expected to contain 1")
	}

	if s.Contains(4) {
		t.Error("Should not contain 4")
	}
}

func TestSet_String(t *testing.T) {
	s := uniqueset.New[string]()
	s.Add("apple")
	s.Add("banana")
	s.Add("apple") // Duplicate!

	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	if !s.Contains("banana") {
		t.Error("Expected to contain 'banana'")
	}
}

func TestSetOperations(t *testing.T) {
	s1 := uniqueset.New[int]()
	s1.AddAll(1, 2, 3)

	s2 := uniqueset.New[int]()
	s2.AddAll(3, 4, 5)

	union := s1.Union(s2)
	if union.Size() != 5 {
		t.Errorf("Union expected size 5, got %d", union.Size())
	}

	intersection := s1.Intersection(s2)
	if !intersection.Contains(3) || intersection.Size() != 1 {
		t.Error("Intersection failed")
	}

	difference := s1.Difference(s2)
	if !difference.Contains(1) || difference.Contains(3) {
		t.Error("Difference failed")
	}
}
