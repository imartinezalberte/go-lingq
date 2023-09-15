package utils

import "github.com/repeale/fp-go"

type (
	Predicate[T any] func(T) bool
	Reduce[T any]    func(T, T) T
)

func SimilarString[U []T, T ~string](input U) []string {
	return fp.Map(func(t T) string {
		return string(t)
	})(input)
}
