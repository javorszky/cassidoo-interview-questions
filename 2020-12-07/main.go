package dec72020

import (
	"fmt"
	"strings"
)

// This week’s question:
// Given two non-negative integers n1 and n2 represented as strings, return the product of n1 and n2, also represented
// as a string. Neither input will start with 0, and don’t just convert it to an integer and do the math that way.
//
// Examples:
//
// $ stringMultiply(“123”, “456”)
// $ “56088”
func Task() {
	fmt.Println("hola")
	stringMultiply("123", "456")
}

func stringMultiply(a, b string) string {
	// matrix (2d slice) with as many rows as len(b), and as many columns as len(a)+1 +len(b)-1, so len(a) + len(b).
	calculator := make([][]string, len(b), len(b))
	for i := range calculator {
		calculator[i] = make([]string, len(a)+len(b), len(a)+len(b))
	}
	fmt.Printf("%#v\n", calculator)

	for c, r := range b {
		fmt.Printf("iterating b: %#v, %#v\n", c, string(r))

	}
	return "0"
}

func reverseString(s string) string {
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteString(string(s[i]))
	}

	return b.String()
}
