package set

import "testing"

func TestNew(t *testing.T) {
	s := New[int]()
	if s.Len() != 0 {
		t.Errorf("Expected empty set, got length %d", s.Len())
	}

	s2 := New(1, 2, 3)
	if s2.Len() != 3 {
		t.Errorf("Expected set length 3, got %d", s2.Len())
	}

	s3 := New(1, 1, 2, 2, 3)
	if s3.Len() != 3 {
		t.Errorf("Expected set length 3 (unique elements only), got %d", s3.Len())
	}
}

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	s.Add(1)
	if s.Len() != 3 {
		t.Errorf("Expected length to remain 3 after adding duplicate, got %d", s.Len())
	}

	if !s.Has(1) || !s.Has(2) || !s.Has(3) {
		t.Error("Set should contain 1, 2, and 3")
	}
}

func TestRemove(t *testing.T) {
	s := New[string]()
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")

	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	s.Remove("banana")
	if s.Len() != 2 {
		t.Errorf("Expected length 2 after removal, got %d", s.Len())
	}

	if s.Has("banana") {
		t.Error("Set should not contain 'banana' after removal")
	}

	s.Remove("nonexistent")
	if s.Len() != 2 {
		t.Errorf("Expected length to remain 2 after removing nonexistent item, got %d", s.Len())
	}
}

func TestHas(t *testing.T) {
	s := New[int]()
	s.Add(42)
	s.Add(100)

	if !s.Has(42) {
		t.Error("Set should contain 42")
	}

	if !s.Has(100) {
		t.Error("Set should contain 100")
	}

	if s.Has(999) {
		t.Error("Set should not contain 999")
	}
}

func TestLen(t *testing.T) {
	s := New[int]()
	if s.Len() != 0 {
		t.Error("Length of empty set should be 0")
	}

	s.Add(1)
	if s.Len() != 1 {
		t.Error("Length should be 1 after adding one item")
	}

	s.Add(1)
	if s.Len() != 1 {
		t.Error("Length should remain 1 after adding duplicate item")
	}

	s.Add(2)
	if s.Len() != 2 {
		t.Error("Length should be 2 after adding another unique item")
	}
}

func TestClear(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.Clear()

	if s.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", s.Len())
	}

	if s.Has(1) {
		t.Error("Set should not contain previously added items after clear")
	}
}

func TestItems(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(1)

	items := s.Items()
	if len(items) != s.Len() {
		t.Errorf("Items length (%d) should match set length (%d)", len(items), s.Len())
	}

	expected := make(map[int]bool)
	expected[1] = true
	expected[2] = true
	expected[3] = true

	for _, item := range items {
		if !expected[item] {
			t.Errorf("Unexpected item %d in Items", item)
		}
		delete(expected, item)
	}

	if len(expected) != 0 {
		t.Errorf("Missing items in Items: %v", expected)
	}
}

func TestGenericTypes(t *testing.T) {
	strSet := New("hello", "world", "hello")
	if strSet.Len() != 2 {
		t.Errorf("String set expected length 2, got %d", strSet.Len())
	}

	boolSet := New(true, false, true)
	if boolSet.Len() != 2 {
		t.Errorf("Bool set expected length 2, got %d", boolSet.Len())
	}

	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}
	p3 := Person{Name: "Alice", Age: 30}

	personSet := New(p1, p2, p3)
	if personSet.Len() != 2 {
		t.Errorf("Person set expected length 2, got %d", personSet.Len())
	}
}
