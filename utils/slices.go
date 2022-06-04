package utils

import (
	"math/rand"
)

// MakeRange generates a sequence of int numbers
func MakeRange(min, max int) []int {
	vals := make([]int, max-min)
	for i := range vals {
		vals[i] = min + i
	}
	return vals
}

// Shuffle input int array using a random number generator
func Shuffle(vals []int, r *rand.Rand) {
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

// ShuffleLen creates a shuffled int array of length
func ShuffleLen(length int, r *rand.Rand) []int {
	vals := MakeRange(0, length)
	Shuffle(vals, r)
	return vals
}

// Sum sequence of integers
func Sum(input ...int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}
