package november292021

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

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

func shift(in []int) (int, []int, error) {
	if len(in) < 1 {
		return 0, nil, errors.New("incoming slice is too short")
	}
	return in[0], in[1:], nil
}
