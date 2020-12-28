package dec282020

import "fmt"

func Task() {
	s := "2220000202220020200"
	fmt.Printf("28th December 2020\n------------------\n\nIndex of '2020' in string '%s' is %d\n", s, find2020full(s))
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

func find2020full(s string) int {
	idx := -1
	if len(s) < 4 {
		return idx
	}

	for i := 0; i < len(s)-3; i++ {
		if s[i:i+4] == "2020" {
			return i
		}
	}

	return idx
}

func find2020byterange(s string) int {
	idx := -1
	two := uint8(50)
	zero := uint8(48)
	if len(s) < 4 {
		return idx
	}
	b := []byte(s)
	for i, c := range b {
		if i > len(s)-3 {
			break
		}
		if c == two {
			if b[i+1] == zero && b[i+2] == two && b[i+3] == zero {
				return i
			}
		}
	}
	return idx
}
