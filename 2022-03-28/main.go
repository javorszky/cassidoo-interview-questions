package march282022

import (
	"errors"
	"fmt"
)

var (
	badStringError     = errors.New("incoming string is not only * or |")
	negativeStartError = errors.New("start index can't be negative")
	outOfBoundsError   = errors.New("end index is past the end of the string")
)

func Tasks() {
	str := "|**|*|*"
	zeroFive, _ := containedItems(str, 0, 5)
	fmt.Printf("%s[%d:%d] has %d items in closed compartments\n", str, 0, 5, zeroFive)

	zeroSix, _ := containedItems(str, 0, 6)
	fmt.Printf("%s[%d:%d] has %d items in closed compartments\n", str, 0, 6, zeroSix)

	oneSeven, _ := containedItems(str, 1, 7)
	fmt.Printf("%s[%d:%d] has %d items in closed compartments\n", str, 1, 7, oneSeven)

	// multi soluiton
	multi, _ := containedItemsMulti(str, [][2]int{
		{0, 5},
		{0, 6},
		{1, 7},
	})
	fmt.Printf("multi indice checks for string [%s] are as follows: %v\n", str, multi)
}

func containedItems(in string, start, end int) (int, error) {
	if start < 0 {
		return 0, negativeStartError
	}

	if end > len(in) {
		return 0, outOfBoundsError
	}

	boundaries, err := parseBoundaries(in)
	if err != nil {
		return 0, err
	}

	blocks, lastBoundary, began := 0, 0, false
	for _, boundary := range boundaries {
		if boundary < start {
			continue
		}

		if boundary > end-1 {
			break
		}

		if began {
			blocks += boundary - lastBoundary - 1
		}
		lastBoundary, began = boundary, true
	}
	return blocks, nil
}

func parseBoundaries(in string) ([]int, error) {
	boundaries := make([]int, 0)
	for i, char := range in {
		switch string(char) {
		case "*":
			continue
		case "|":
			boundaries = append(boundaries, i)
		default:
			return nil, badStringError
		}
	}

	return boundaries, nil
}

func containedItemsMulti(in string, boundaries [][2]int) ([]int, error) {
	stringBoundaries, err := parseBoundaries(in)
	if err != nil {
		return nil, err
	}

	results := make([]int, len(boundaries))

	for i, edges := range boundaries {
		if edges[0] < 0 {
			return nil, negativeStartError
		}

		if edges[1] > len(in) {
			return nil, outOfBoundsError
		}

		blocks, lastBoundary, began := 0, 0, false
		for _, localBoundary := range stringBoundaries {
			if localBoundary < edges[0] {
				continue
			}

			if localBoundary > edges[1]-1 {
				break
			}

			if began {
				blocks += localBoundary - lastBoundary - 1
			}
			lastBoundary, began = localBoundary, true
		}

		results[i] = blocks
	}

	return results, nil
}
