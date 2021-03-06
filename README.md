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
	"time"

	"github.com/danielnegri/random-go"
)

func main() {
	probabilities := []float64{0.35, 0.20, 0.40, 0.05}
	random.Seed(time.Now().UnixNano()) // Optional

	// Select a weighted random selected
	index := random.SelectWeightedIndex(probabilities)
	fmt.Printf("Index: %d", index)
}
```
