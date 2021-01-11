package sep142020

import (
	"fmt"
	"math/rand"
)

func Task() {
	integers := []int{1, 3, 7, 11, 12, 14, 18}

	length := fibonacciLike(integers)
	fmt.Printf("14th September 2020\n-------------------\n")
	fmt.Printf("Longest sequence is %d long in the original sequence\n", length)

	n := 100
	newList := generateList(n)
	newLength := fibonacciLike(newList)
	fmt.Printf("Longest sequence is %d long in a 100 element random monotonic increasing list of integers\n", newLength)

	fiboList := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 90}
	fiboLength := fibonacciLike(fiboList)
	fmt.Printf("Longest sequence is %d long in an 11 element mostly fibonacci sequence\n\n", fiboLength)
}

func fibonacciLike(integers []int) int {
	var (
		thirdNumber int
	)
	m := make(map[int]struct{})
	current := make([]int, 0)
	lists := make([][]int, 0)

	// Create a map for efficient lookup.
	for _, i := range integers {
		m[i] = struct{}{}
	}

	for firstIndex, firstNumber := range integers {
		for secondIndex, secondNumber := range integers {
			// We don't want to go backwards.
			if secondIndex <= firstIndex {
				continue
			}

			l := len(current)

			// If we're just starting a new check, populate the first two numbers and move on.
			if l < 2 {
				current = append(current, firstNumber, secondNumber)
			}

			l = len(current)
			thirdNumber = current[l-1] + current[l-2]

			_, ok := m[thirdNumber]
			if !ok {
				// The sequence ends, add current sequence to the list of sequences if longer than 2, truncate the
				// current sequence, and move on to the next one.
				if l > 2 {
					lists = append(lists, current)
				}
				current = []int{}
				continue
			}

			current = append(current, thirdNumber)
		}
	}

	// If it's the last loop, check whether we have a long enough sequence, if so, add it to the list, and clean up.
	if len(current) > 2 {
		lists = append(lists, current)
	}
	current = []int{}

	maxLength := 0
	for _, seq := range lists {
		if len(seq) > maxLength {
			maxLength = len(seq)
		}
	}

	return maxLength
}

func generateList(n int) []int {
	l := make([]int, 0)
	previous := 0
	for i := 0; i < n; i++ {
		previous = previous + rand.Intn(10)
		l = append(l, previous)
	}

	return l
}
