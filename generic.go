package main

// Define an interface that allows ordering operations
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

// Now we can use the > operator with this constraint
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item, true
}
