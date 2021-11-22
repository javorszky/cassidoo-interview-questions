package november222021

import (
	"fmt"
	"sort"
	"strings"
)

var letterCodeToPrime = map[int32]int{
	97:  2,
	98:  3,
	99:  5,
	100: 7,
	101: 11,
	102: 13,
	103: 17,
	104: 19,
	105: 23,
	106: 29,
	107: 31,
	108: 37,
	109: 41,
	110: 43,
	111: 47,
	112: 53,
	113: 59,
	114: 61,
	115: 67,
	116: 71,
	117: 73,
	118: 79,
	119: 83,
	120: 89,
	121: 97,
	122: 101,
}

func Tasks() {
	newsletter := []string{"eat", "tea", "ten", "poop", "net", "ate"}

	fmt.Printf("22nd November 2021\n"+
		"------------------\n"+
		"After grouping the follwing input:\n"+
		"%% %v\n\n"+
		""+
		"The grouped output by anagrams become:\n"+
		"%% %v\n", newsletter, groupByAnagrams(newsletter))
}

func groupByAnagrams(listOfStuff []string) [][]string {
	anagramGroups := make([][]string, 0)
	groupedByNumbers := make(map[int][]string)

	for _, word := range listOfStuff {
		start := 1
		for _, letter := range word {
			start = start * letterCodeToPrime[letter]
		}

		groupedByNumbers[start] = append(groupedByNumbers[start], word)
	}

	for _, words := range groupedByNumbers {
		anagramGroups = append(anagramGroups, words)
	}
	return anagramGroups
}

func groupByAnagramsButtStrings(listOfStuff []string) [][]string {
	anagramGroups := make([][]string, 0)
	groupedBySortedLetters := make(map[string][]string)

	for _, word := range listOfStuff {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		key := strings.Join(letters, "")
		if groupedBySortedLetters[key] == nil {
			groupedBySortedLetters[key] = make([]string, 0)
		}
		groupedBySortedLetters[key] = append(groupedBySortedLetters[key], word)
	}

	for _, words := range groupedBySortedLetters {
		anagramGroups = append(anagramGroups, words)
	}
	return anagramGroups
}
