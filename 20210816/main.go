package august162021

import (
	"errors"
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
	var fEven func(string, int, int) string

	wrap := func(what string, idxDown, idxUp int) (string, error) {
		// next step would be out of bounds, so whatever we have now is a palindrome that touches at least one edge.
		if idxDown < 0 || idxUp >= len(stringSlice) {
			return what, errors.New("next is out of bounds")
		}

		// next steps are not the same letter, so it would stop being a palindrome.
		if stringSlice[idxDown] != stringSlice[idxUp] {
			return what, errors.New("next is not palindrome")
		}

		return stringSlice[idxDown] + what + stringSlice[idxUp], nil
	}

	f = func(center string, origin, distance int) string {
		down := origin - distance
		up := origin + distance

		center, err := wrap(center, down, up)
		if err != nil {
			return center
		}

		return f(center, origin, distance+1)
	}

	fEven = func(center string, origin, distance int) string {
		down := origin - distance + 1
		up := origin + distance

		center, err := wrap(center, down, up)
		if err != nil {
			return center
		}

		return fEven(center, origin, distance+1)
	}

	for i := 1; i < len(stringSlice); i++ {
		try := f(stringSlice[i], i, 1)
		tryEven := fEven("", i, 1)
		if len(try) > len(longest) {
			longest = try
		}

		if len(tryEven) > len(longest) {
			longest = tryEven
		}
	}

	return longest
}
