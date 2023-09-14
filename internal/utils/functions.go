package utils

import (
	"encoding/json"
	"io"

	"golang.org/x/exp/constraints"
)

func UniqueValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// JSON Unmarshals the given ReadCloser buffer into target ensuring that
// the buffer is read until EOF.
func JSONUnmarshalTilEOF(r io.ReadCloser, target interface{}) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

// RemoveDuplicates is a simple function to remove duplicates from an ordered array.
//
//	input := []PersonalString{"hello", "bye", "hello"}
//
//	fmt.Println(RemoveDuplicates(input)) // output hello and bye
func RemoveDuplicates[T constraints.Ordered](input []T) []T {
	var (
		helper = make(map[T]interface{})
		output []T
	)

	for i := range input {
		if _, ok := helper[input[i]]; !ok {
			helper[input[i]] = 0
			output = append(output, input[i])
		}
	}

	return output
}

// RemoveDuplicatesKeepingLastSeen function parses an array of any structure and returns another array
// removing duplicates keeping inside of it the first ones.
// mapFunc allows us to get a key from the structure, so we can create the map with some kind of order.
// the key must be an ordered value.
//
//	type (
//		PersonalString string
//		Person struct {
//			ID PersonalString
//			Name string
//			Age int
//		}
//	)
//
//	input := []Person{
//		{ ID: "1", Name: "Pedro", Age: 18 },
//		{ ID: "2", Name: "Juan", Age: 30 },
//		{ ID: "3", Name: "Diego", Age: 35 },
//		{ ID: "2", Name: "Ander", Age: 24 },
//	}
//
//	output := RemoveDuplicatesKeepingFirstSeen(input, func(person Person) PersonalString {
//		return person.ID
//	})
//
//	fmt.Println(output) // contains Pedro, Diego and Juan
func RemoveDuplicatesKeepingFirstSeen[K constraints.Ordered, V any](
	input []V,
	mapFunc func(V) K,
) []V {
	var (
		helper = make(map[K]V)
		output []V
	)

	for i := range input {
		key := mapFunc(input[i])
		if _, ok := helper[key]; !ok {
			helper[key] = input[i]
		}
	}

	for _, val := range helper {
		output = append(output, val)
	}

	return output
}

// RemoveDuplicatesKeepingLastSeen function parses an array of any structure and returns another array
// removing duplicates keeping inside of it the last ones.
// mapFunc allows us to get a key from the structure, so we can create the map with some kind of order.
// the key must be an ordered value.
//
//	type (
//		// This is used just to show that we don't need a strict int, string, float or complex type to used as a key.
//		PersonalString string
//		Person struct {
//			ID PersonalString
//			Name string
//			Age int
//		}
//	)
//
//	input := []Person{
//		{ ID: "1", Name: "Pedro", Age: 18 },
//		{ ID: "2", Name: "Juan", Age: 30 },
//		{ ID: "3", Name: "Diego", Age: 35 },
//		{ ID: "2", Name: "Ander", Age: 24 },
//	}
//
//	output := RemoveDuplicatesKeepingLastSeen(input, func(person Person) PersonalString {
//		return person.ID
//	})
//
//	fmt.Println(output) // contains Pedro, Diego and Ander
func RemoveDuplicatesKeepingLastSeen[K constraints.Ordered, V any](
	input []V,
	mapFunc func(V) K,
) []V {
	var (
		helper = make(map[K]V)
		output []V
	)

	for i := range input {
		helper[mapFunc(input[i])] = input[i]
	}

	for _, val := range helper {
		output = append(output, val)
	}

	return output
}

// Returns a if the evaluated condition cond is truthy. Returns b instead.
// Useful when conditionally get a certain value in one line.
//
//	age := 18
//	// Returns "coke" because condition is false
//	drink := IfElse(age > 21, "beer", "coke")
func IfElse[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}

// Returns a function with the condition preloaded, so you don't have to call it over and over again.
//
//	ifEmptyOrElse := IfElseCond(func(a string) bool {
//		return strings.TrimSpace(a) != ""
//	})
//
//	Expect(ifEmptyOrElse("", "second")).To(Equal("second"))
//	Expect(ifEmptyOrElse("first", "second")).To(Equal("first"))
func IfElseCond[T any](cond Predicate[T]) Reduce[T] {
	return func(a T, b T) T {
		if cond(a) {
			return a
		}
		return b
	}
}

// MergeMaps function merges two maps moving the src keys and values to the dst ones, without overriden them.
//
//	import "fmt"
//
//	src, dst := map[string]string { "a": "Hello", "b": "Bye" }, map[string]string { "a": "Salut", "c": "Aurevoir" }
//
//	MergeMaps(src, dst)
//
//	fmt.Println(dst) // Prints: map[a:Salut b:Bye c:Aurevoir]
func MergeMaps[T comparable, K any](src map[T]K, dst map[T]K) map[T]K {
	for k, v := range src {
		if _, ok := dst[k]; !ok {
			dst[k] = v
		}
	}
	return dst
}

func GetOrPanic[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}

	return res
}

func GetOrPanicF[T any](f func() (T, error)) func() T {
	return func() T {
		return GetOrPanic(f())
	}
}
