package may162022

import (
	"errors"
	"fmt"
	"log"
)

func Tasks() {
	slice1 := []int{7, 4, 10, 0, 1}
	slice2 := []int{9, 7, 2, 3, 6}

	maxiSlice, err := maximizeSlices(slice1, slice2)
	if err != nil {
		log.Fatalf("encountered error: %s", err.Error())
	}

	fmt.Printf("Maximised slice is %v\n", maxiSlice)
}

func maximizeSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) != len(slice2) {
		return nil, errors.New("the two slices are not the same length")
	}

	buf := 0
	bigger, smaller := 0, 0
	result := make(map[int]struct{})

	for i := range slice1 {
		a, b := slice1[i], slice2[i]

		fmt.Printf("\n"+
			"*--------------*\n"+
			"| Iteration %2d |\n"+
			"*--------------*\n"+
			"\nSo far\n%v\n\n%d vs %d with buf %d\n", i, result, a, b, buf)
		bigger = b
		smaller = a

		if a > b {
			bigger = a
			smaller = b
		}

		if buf != 0 && buf > bigger {
			if _, ok := result[buf]; !ok {
				fmt.Printf(" - buf (%d) was not 0, was bigger than the bigger number (%d), it got stored, reset to the bigger of the other two numbers that aren't in the map yet, or 0\n", buf, bigger)
				result[buf] = struct{}{}
				if _, ok := result[bigger]; !ok {
					buf = bigger
					continue
				}

				if _, ok := result[smaller]; !ok {
					buf = smaller
					continue
				}

				buf = 0
				continue
			}
		}

		if _, ok := result[bigger]; !ok {
			fmt.Printf(" - buf (%d) wasn't good (either 0 or not bigger than bigger), bigger (%d) doesn't exist in the map yet , it got stored\n", buf, bigger)
			result[bigger] = struct{}{}
			if buf < smaller {
				fmt.Printf(" - buf (%d) was smaller than the smaller number (%d), so that got replaced, continuing\n", buf, smaller)
				buf = smaller
				continue
			}

			if _, ok = result[buf]; ok {
				fmt.Printf(" - buf (%d) was equal or bigger than the smaller number (%d), but buf is already in the map, so replacing buf with smaller anyways\n", buf, smaller)
				buf = smaller
			}

			continue
		}

		if _, ok := result[smaller]; !ok {
			fmt.Printf(" - bigger (%d) was in the map, smaller (%d) isn't, storing that, resetting buf to 0\n", bigger, smaller)
			result[smaller] = struct{}{}
			buf = 0
			continue
		}

		return nil, errors.New("encountered an impossible situation")
	}

	rslice := make([]int, 0, len(result))
	for k := range result {
		rslice = append(rslice, k)
	}

	return rslice, nil
}
