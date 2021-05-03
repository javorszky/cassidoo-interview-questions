package may032021

import (
	"fmt"
	"math"
)

func Task() {
	nums := []float64{
		1.7,
		-2.1,
		500.4,
		-369.5,
		150,
		-350,
	}
	for _, num := range nums {
		fmt.Printf("rounding %f to nearest integer closer to 0 is %f\n", num, round(num))
	}
}

// round will "round" the number to the next integer closer to zero.
func round(n float64) float64 {
	return math.Trunc(n)
}
