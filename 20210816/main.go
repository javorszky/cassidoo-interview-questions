package august162021

import (
	"fmt"
	"strings"
)

func Tasks() {
	const inc = "badad"
	palindrome := pSubString(inc)

	fmt.Printf("The first longest palindrome in '%s' is '%s'\n", inc, palindrome)
}

func pSubString(incoming string) string {
	stringSlice := strings.Split(incoming, "")
	longest := stringSlice[0]

	var f func(string, int, int) string

	f = func(center string, origin, distance int) string {
		up := origin + distance
		down := origin - distance

		// next step would be out of bounds, so whatever we have now is a palindrome that touches at least one edge.
		if down < 0 || up >= len(stringSlice) {
			return center
		}

		// next steps are not the same letter, so it would stop being a palindrome.
		if stringSlice[down] != stringSlice[up] {
			return center
		}

		center = stringSlice[down] + center + stringSlice[up]

		return f(center, origin, distance+1)
	}

	for i := 1; i < len(stringSlice); i++ {
		try := f(stringSlice[i], i, 1)
		if len(try) > len(longest) {
			longest = try
		}
	}

	return longest
}
