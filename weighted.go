package random

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectWeightedIndex returns a weighted random index given a list of probabilities.
// For example, given [2, 3, 5] it returns 0 (the index of the first element) with
// probability 0.2, 1 with probability 0.3 and 2 with probability 0.5. The weights
// need not sum up to anything in particular, and can actually be arbitrary floating
// point numbers.
//
// Time Complexity: O(N)
// Space Complexity: O(N)
func SelectWeightedIndex(weights []float64, seed int64) (int, error) {
	if len(weights) == 0 {
		return 0, fmt.Errorf("invalid parameters: weights cannot be empty")
	}

	var totals = make([]float64, len(weights))
	var total float64

	for i, choice := range weights {
		total += choice
		totals[i] = total
	}

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	rand.Seed(seed)
	rnd := rand.Float64() * total
	for i, v := range totals {
		if rnd < v {
			return i, nil
		}
	}

	return 0, nil
}
