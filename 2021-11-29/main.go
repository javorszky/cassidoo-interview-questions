package november292021

import (
	"fmt"
	"regexp"
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

	return numberMap[2], nil
}
