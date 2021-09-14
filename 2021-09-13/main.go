package september132021

import (
	"fmt"
	"strings"
)

const pairsOfParens = 5

type parens []string

func (p parens) open() parens {
	return append(p, "(")
}

func (p parens) close() parens {
	return append(p, ")")
}

func (p parens) String() string {
	var sb strings.Builder
	for _, i := range p {
		sb.WriteString(i)
	}
	return sb.String()
}

func (p parens) Hill() string {
	elevation := make(map[int]map[int]string)
	e := 0
	for i, v := range p {
		switch v {
		case "(":
			e = e + 1
			if elevation[e] == nil {
				elevation[e] = make(map[int]string)
			}
			elevation[e][i] = v
		case ")":
			elevation[e][i] = v
			e = e - 1
		default:
			panic("// this will never happen")
		}
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("o%so\n", strings.Repeat("-", len(p))))
	maxHeight := len(p) / 2
	for j := maxHeight; j > 0; j-- {
		sb.WriteString("|")
		relief, ok := elevation[j]
		if !ok {
			sb.WriteString(strings.Repeat(" ", len(p)) + "|\n")
			continue
		}
		for k := 0; k < len(p); k++ {
			c, ok := relief[k]
			if !ok {
				sb.WriteString(" ")
			} else {
				sb.WriteString(c)
			}
		}
		sb.WriteString("|\n")
	}
	sb.WriteString(fmt.Sprintf("o%so", strings.Repeat("-", len(p))))
	return sb.String()
}

func Tasks() {
	ps := constructParens(pairsOfParens)
	fmt.Printf("13th September 2021: all possible paren configs with %d pairs of parens:\n", pairsOfParens)
	for _, ppair := range ps {
		fmt.Printf("%s\n", ppair.Hill())
	}

	ps[12].Hill()

}

func constructParens(n int) []parens {
	var f func(parens, int, int, int) []parens

	f = func(p parens, value, left, step int) []parens {
		acc := make([]parens, 0)

		if step == 0 {
			return append(acc, p)
		}

		step = step - 1

		if left > 0 {
			acc = append(acc, f(p.open(), value+1, left-1, step)...)
		}

		if value > 0 {
			acc = append(acc, f(p.close(), value-1, left, step)...)
		}

		return acc
	}

	p := parens{}

	return f(p.open(), 1, n-1, n*2-1)
}
