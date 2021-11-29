package november292021

import (
	"fmt"
)

func Tasks() {
	input := "23"

	output, _ := phoneLetter(input)

	fmt.Printf("29th November 2021\n"+
		"------------------\n"+
		"Possible letter combinations for input [%s]\n"+
		"%v\n\n", input, output)
}

func phoneLetter(s string) ([]string, error) {
	if !numRegex.MatchString(s) {
		return nil, fmt.Errorf("incoming string is not numbers only, has numbers [0, 1] in them, or is empty: %s", s)
	}

	numbers, err := parseNumbers(s)
	if err != nil {
		return nil, fmt.Errorf("parsenumbers failed for some weird reason. This should not have happened: %w", err)
	}

	return assembler("", numbers)
}

func assembler(prefix string, numbers []int) ([]string, error) {
	number, rest, err := shift(numbers)
	if err != nil {
		return nil, fmt.Errorf("assembler: shift returned error: %w", err)
	}
	letters, ok := numberMap[number]
	if !ok {
		panic(fmt.Sprintf("fetching the list of letters from the map for number %d failed. This is a catastrophic failure!", number))
	}

	branch := make([]string, 0)
	for _, l := range letters {
		eh, err := assembler(prefix+l, rest)
		if err != nil {
			branch = append(branch, prefix+l)
			continue
		}
		branch = append(branch, eh...)
	}

	return branch, nil
}
