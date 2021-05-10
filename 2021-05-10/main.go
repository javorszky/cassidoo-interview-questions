package may102021

import (
	"fmt"
	"strconv"

	dec72020 "github.com/javorszky/cassidoo-interview-questions/2020-12-07"
)

const (
	checkInteger = 1234567890
	checkString  = "111111112222222233333333"
)

func Task() {
	fmt.Printf("the digits of %d are the same when cubed: %t\n", checkInteger, sameDigits(checkInteger))
	fmt.Printf("the digits of %s are the same when cubed: %t\n", checkString, sameDigitsString(checkString))
}

func sameDigits(n int) bool {
	return sameDigitsString(strconv.Itoa(n))
}

func sameDigitsString(n string) bool {
	n2String := dec72020.StringMultiply(n, n)
	n3String := dec72020.StringMultiply(n2String, n)

	set := make(map[string]struct{})
	for _, s := range n {
		set[string(s)] = struct{}{}
	}

	for _, s := range n3String {
		_, ok := set[string(s)]
		if !ok {
			return false
		}
	}
	return true
}
