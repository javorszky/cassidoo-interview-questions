package oct52020

import (
	"fmt"
	"math/rand"
	"sort"
)

func Task() {
	unsortedSlice := []int{3, 8, 1, 4, 6, 3, 7, 1, 2, 9, 7, 2, 6, 7, 8, 2, 15, 4, 3}

	pivot := rand.Intn(len(unsortedSlice))

	spSlice := sortAndRotateSlice(unsortedSlice, pivot)

	lowest := lowestValue(spSlice)

	fmt.Printf("5th October 2020\n----------------\nThe lowest value of the previously sorted and rotated slice of\n%v\n is %d\n\n", spSlice, lowest)
}

// takes a slice of integers that have been previously sorted and pivoted, and returns the lowest value from it in O(n)
// time or less.
func lowestValue(integers []int) int {
	// if the slice is empty, return 0. Sensible default.
	if len(integers) == 0 {
		return 0
	}

	// start with assuming the first element in the slice is the lowest.
	lowest, check := integers[0], integers[0]
	for _, i := range integers {
		if i < check {
			// the moment we have a lower number than the previous, return the lower number, because it's a sorted slice
			// and we have found the pivot point.
			return i
		}
		// the new number i was not smaller than the one we looked at, so advance the pointer and try again.
		check = i
	}
	// if we've not returned by this point, that means that the slice was pivoted at the beginning or the end, or the
	// elements are the same, and in all cases the first element of the slice is actually the lowest value.
	return lowest
}

// sortAndRotateSlice takes a slice of integers, makes sure it's sorted in ascending order, and then rotates it at some
// point to ensure the start condition of our algorithm is satisfied.
func sortAndRotateSlice(integers []int, pivot int) []int {
	sort.Ints(integers)

	return append(integers[pivot:], integers[:pivot]...)
}
