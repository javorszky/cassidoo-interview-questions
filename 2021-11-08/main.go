package november082021

import "fmt"

func Tasks() {
	t1 := []int{1, 2, 3, 1}
	t2 := []int{1, 3, 2, 3, 5, 6, 4}

	fmt.Println("8th november 2021\n-----------------")
	for _, v := range [][]int{t1, t2} {
		fmt.Printf("The local peaks for array %v are\n%v\n\n", v, localPeaks(v))
	}
}

func localPeaks(stream []int) []int {
	previous := 0
	peaks := make([]int, 0)
	plateau := make([]int, 0)

	// for every next number in the stream, check if this number is smaller. If it is smaller, append the contents of
	// the index plateaus to the peaks slice.
	for i, v := range stream {
		// If this is the very first number in the array, record it into previous, but ignore anything else. Because
		// there's no previous value, this can't be a peak because it can't be higher than a non existent number.
		if i == 0 {
			previous = v
			continue
		}

		// If current number is smaller than the previous that means we're either closing a peak, or we're on a descent.
		// In either case we should append what's in the plateau slice to the permanent record. If we're on a descent
		// this will append an empty slice to a slice, not changing anything about it. It's easier to do this than to
		// code in extra checks for what the plateau slice might be.
		//
		// Also store the current value into previous, so we can find a local peak, and skip to next.
		if v < previous {
			peaks = append(peaks, plateau...)
			plateau = make([]int, 0)
			previous = v
			continue
		}

		// If the current value is larger than the previous one, that means we're either on a new peak, or on an ascent.
		// In either case, make the plateau slice contain _only_ the index of the current element, ignoring a previous
		// potential value of the slice from the previous number, which might also have been a suspected peak / ascent.
		//
		// Store current as previous, and skip to next.
		if v > previous {
			plateau = []int{i}
			previous = v
			continue
		}

		// If we're here, the value of the current number is the same as the previous, so store the current index in the
		// running plateau index for further collection.
		plateau = append(plateau, i)
	}

	return peaks
}
