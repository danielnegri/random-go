package main

import (
	"fmt"

	"github.com/danielnegri/random-go"
)

func main() {
	probabilities := []float64{0.35, 0.20, 0.40, 0.05}

	// Select a weighted random selected
	index := random.SelectWeightedIndex(probabilities)
	fmt.Printf("Index: %d", index)
}
