package sept262022

import (
	"fmt"
)

const (
	th = "th"
	st = "st"
	nd = "nd"
	rd = "rd"
)

func Tasks() {
	fmt.Println("september 26th, 2022 solution")
	fmt.Println(ordinal(3))
	fmt.Println(ordinal(57))
}

func ordinal(n uint64) string {
	ones := findOnes(n)
	tens := findTens(n)
	switch {
	case tens == 1:
		return stringify(n, th)
	case ones == 1:
		return stringify(n, st)
	case ones == 2:
		return stringify(n, nd)
	case ones == 3:
		return stringify(n, rd)
	default:
		return stringify(n, th)
	}
}

func findOnes(n uint64) uint64 {
	return n % 10
}

func findTens(n uint64) uint64 {
	return findOnes(n / 10)
}

func stringify(n uint64, suffix string) string {
	return fmt.Sprintf("%d%s", n, suffix)
}
