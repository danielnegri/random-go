# Random (GO)

Random is a collection of useful methods for random selections in Go.

## Usage

Download this package.

```
$ go get github.com/danielnegri/random-go
```

To import this package, add the following line to your code:

```
import "github.com/danielnegri/random-go"
```


## Sample

The following is a basic testing example.

Main program:

```go
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
```
