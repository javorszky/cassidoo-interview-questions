package november292021

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Tasks() {

}

var numRegex = regexp.MustCompile(`^[2-9]+$`)
var numberMap = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

func phoneLetter(s string) ([]string, error) {
	if !numRegex.MatchString(s) {
		return nil, fmt.Errorf("incoming string is not numbers only, has numbers [0, 1] in them, or is empty: %s", s)
	}

	numbers, err := parseNumbers(s)
	if err != nil {
		return nil, fmt.Errorf("parsenumbers failed for some weird reason. This should not have happened: %w", err)
	}

	return assembler("", numbers), nil
}

func parseNumbers(s string) ([]int, error) {
	nums := make([]int, len(s))
	for i, l := range s {
		n, err := strconv.Atoi(string(l))
		if err != nil {
			return nil, fmt.Errorf("parseNumbers: strconv.Atoi for %s: %w", string(l), err)
		}
		nums[i] = n
	}
	return nums, nil
}

func assembler(prefix string, numbers []int) []string {

	return nil
}

func shift(in []int) (int, []int, error) {
	if len(in) < 1 {
		return 0, nil, errors.New("incoming slice is too short")
	}
	return in[0], in[1:], nil
}
