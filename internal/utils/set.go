package utils

type Set[T comparable] map[T]any

func (s Set[T]) Exists(input T) bool {
	_, ok := s[input]
	return ok
}

func (s Set[T]) ToArr() []T {
	arr := make([]T, 0)
	for k := range s {
		arr = append(arr, k)
	}
	return arr
}

func (s Set[T]) Add(input T) {
	if !s.Exists(input) {
		s[input] = nil
	}
}

func NewSet[T comparable]() Set[T] {
	return make(map[T]any)
}
