package nov212022

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

const (
	// backslash is this one \
	backslash int32 = 92
	// slash is this one /
	slash int32 = 47
)

func Tasks() {
	slashes := []string{
		`\\\\`,
		`\\\//\/\\`,
		`//\/\\/\///\\\\/\/\/\\/\\\//\\/\/\\/\/`,
	}

	for _, sl := range slashes {
		fmt.Printf("\nPrinting vertical for %s\n\n", sl)
		err := printSlashes(sl)
		if err != nil {
			log.Fatalf("encountered issue with printSlashes for the following string:\n%s\n\nerr: %s", sl, err.Error())
		}
	}

}

func printSlashes(in string) error {
	nums, err := convertSlashes(in)
	if err != nil {
		return errors.Wrap(err, "convertSlashes")
	}

	for i, n := range nums {
		pad := strings.Repeat(" ", n)
		switch in[i] {
		case uint8(slash):
			pad += " "
		}
		fmt.Printf("%s%s\n", pad, string(in[i]))
	}

	return nil
}

func convertSlashes(in string) ([]int, error) {
	out := make([]int, len(in))
	start := 0
	min := 0
	for i, c := range in {
		switch c {
		case slash:
			start--
		case backslash:
			start++
		default:
			return nil, fmt.Errorf("invalid character %d, %s", c, string(c))
		}
		if start < min {
			min = start
		}

		out[i] = start
	}

	// normalize the slice
	for i, n := range out {
		out[i] = n - min
	}

	return out, nil
}
