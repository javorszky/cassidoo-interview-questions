package august302021

import (
	"fmt"
	"strings"
)

//  Write a function to find the longest common prefix string in an array of strings.

func Tasks() {
	sts := []string{"cranberry", "crawfish", "crap"}
	fmt.Printf("longest prefix for %v is '%s'\n", sts, longestPrefix(sts))
}

type graph map[int]map[string]struct{}

func longestPrefix(arrayOfStrings []string) string {
	maxLen := 0
	for _, s := range arrayOfStrings {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	g := make(graph)

	for _, st := range arrayOfStrings {
		letters := strings.Split(st, "")
		for i, letter := range letters {
			if g[i] == nil {
				g[i] = make(map[string]struct{})
			}

			g[i][letter] = struct{}{}
		}
	}

	var sb strings.Builder
	for j := 0; j < maxLen; j++ {
		if len(g[j]) != 1 {
			return sb.String()
		}

		for k := range g[j] {
			sb.WriteString(k)
		}
	}

	return sb.String()
}
