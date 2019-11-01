package main

import (
	"fmt"
	"github.com/danielnegri/random-go"
)

func main() {
	probabilities := []float64{0.35, 0.20, 0.40, 0.05}

	// Select a weighted random selected
	index, err := random.SelectWeightedIndex(probabilities, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Index: %d", index)
}
