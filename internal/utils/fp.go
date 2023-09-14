package utils

type (
	Predicate[T any] func(T) bool
	Reduce[T any]    func(T, T) T
)
