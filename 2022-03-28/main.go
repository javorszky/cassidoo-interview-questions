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
	fmt.Printf("%s[%d:%d] has %d items in closed compartments\n", str, 0, 5)
	fmt.Printf("%s[%d:%d] has %d items in closed compartments\n", str, 0, 6)
	fmt.Printf("%s[%d:%d] has %d items in closed compartments\n", str, 1, 7)
}

func containedItems(in string, start, end int) (int, error) {
	if start < 0 {
		return 0, negativeStartError
	}

	if end > len(in) {
		return 0, outOfBoundsError
	}

	began, items, closedItems := false, 0, 0

	for _, char := range in[start:end] {
		switch string(char) {
		case "*":
			if !began {
				continue
			}
			items++
		case "|":
			began = true
			closedItems += items
			items = 0
		default:
			return 0, badStringError
		}
	}

	return closedItems, nil
}
