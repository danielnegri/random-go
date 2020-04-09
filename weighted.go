package random

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Seed(s int64) {
	random = rand.New(rand.NewSource(s))
}

// SelectWeightedIndex returns a weighted random index given a list of probabilities.
// For example, given [2, 3, 5] it returns 0 (the index of the first element) with
// probability 0.2, 1 with probability 0.3 and 2 with probability 0.5. The weights
// need not sum up to anything in particular, and can actually be arbitrary floating
// point numbers.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func SelectWeightedIndex(weights []float64) int {
	var totals = make([]float64, len(weights))
	var total float64

	for i, choice := range weights {
		total += choice
		totals[i] = total
	}

	rnd := random.Float64() * total
	for i, v := range totals {
		if rnd < v {
			return i
		}
	}

	return 0
}
