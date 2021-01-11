package aug242020

import (
	"fmt"
	"strings"
)

func Task() {
	s := "oh heavens"
	c := "h"
	num := numChars(s, c)
	mNum := numCharsManual(s, c)

	fmt.Print("24th August 2020\n----------------\n")
	fmt.Printf("library: the string '%s' contains '%s' %d times\n", s, c, num)
	fmt.Printf("manual: the string '%s' contains '%s' %d times\n\n", s, c, mNum)
}

func numChars(s, c string) int {
	return strings.Count(s, c)
}

func numCharsManual(s, c string) int {
	letters := strings.Split(s, "")
	count := 0
	for _, v := range letters {
		if v == c {
			count++
		}
	}
	return count
}
