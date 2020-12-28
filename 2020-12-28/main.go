package dec282020

import "fmt"

func Task() {
	s := "2220000202220020200"
	fmt.Printf("28th December 2020\n------------------\n\nIndex of '2020' in string '%s' is %d\n", s, find2020(s))
}

// You’re given a string of characters that are only 2s and 0s. Return the index of the first occurrence of “2020”
// without using the indexOf (or similar) function, and -1 if it’s not found in the string.

func find2020(s string) int {
	found := -1
	if len(s) < 4 {
		return found
	}

	for i := 0; i < len(s)-3; i++ {
		switch s[i : i+1] {
		case "2":
			if s[i+1:i+4] == "020" {
				return i
			}
		}
	}

	return found
}
