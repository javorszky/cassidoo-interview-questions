package dec72020

import (
	"fmt"
	"strings"
)

const (
	smolA = "123"
	smolB = "456"
	bigA  = "982379879879874937985893274598347893457893475"
	bigB  = "11736218823949872398237492387493278947923794723749238"
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
	fmt.Printf("7th December 2020\n-----------------\n\n")
	result := stringMultiply(smolA, smolB)
	fmt.Printf("the product of these two numbers using only string methods is\n  %s\n* %s\n%s\n  %s\n\n\n", smolA, smolB, strings.Repeat("-", len(smolA)), result)

	res2 := stringMultiply(bigA, bigB)
	fmt.Printf("the product of these two numbers using only string methods is\n  %s\n* %s\n%s\n  %s\n\n", bigA, bigB, strings.Repeat("-", len(bigB)), res2)
}

func stringMultiply(a, b string) string {
	// matrix (2d slice) with as many rows as len(b), and as many columns as len(a)+1 +len(b)-1, so len(a) + len(b).
	calculator := make([][]string, len(b), len(b))
	for i := range calculator {
		calculator[i] = make([]string, len(a)+len(b), len(a)+len(b))
	}

	for c, r := range reverseString(b) {
		overflow := "0"
		var idx int
		for d, s := range reverseString(a) {
			product := multiplication[string(s)][string(r)]
			preOverflow := "0"
			if len(product) > 1 {
				preOverflow = string(product[0])
				product = string(product[1])
			}
			product = sum[product][overflow]
			overflow = preOverflow
			preOverflow = "0"
			if len(product) > 1 {
				overflow = sum[overflow][string(product[0])]
				product = string(product[1])
			}
			calculator[c][c+d] = product
			idx = d
		}
		if overflow != "0" {
			calculator[c][c+idx+1] = overflow
		}
	}

	rotated := rotateMatrix(calculator)
	var result strings.Builder
	overflow := make(map[int]string, 0)

	for ridx, row := range rotated {
		add := addNumbers(row...)

		if c, ok := overflow[ridx]; ok {
			add = addNumbers(add, c)
		}
		add = reverseString(add)

		if len(add) == 1 {
			result.WriteString(add)
			continue
		}
		result.WriteString(string(add[0]))
		for aidx, addDigit := range add[1:] {
			overflow[ridx+1+aidx] = addNumbers(overflow[ridx+1+aidx], string(addDigit))
		}
	}

	return strings.TrimLeft(reverseString(result.String()), "0")
}

func reverseString(s string) string {
	var b strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteString(string(s[i]))
	}

	return b.String()
}

func addNumbers(a ...string) string {
	// find out how wide our matrix needs to be.
	maxLen := 0
	for _, n := range a {
		if len(n) > maxLen {
			maxLen = len(n)
		}
	}

	if maxLen == 0 {
		return "0"
	}

	// create the empty matrix with len(a) tall and maxLen wide.
	matrix := make([][]string, len(a))
	for i := range matrix {
		matrix[i] = make([]string, maxLen)
	}

	// fill the matrix with the reversed numbers from left to right.
	for i, n := range a {
		for j, letter := range reverseString(n) {
			matrix[i][j] = string(letter)
		}
	}

	// rotate the matrix so we can deal with addition row by row
	rotated := rotateMatrix(matrix)

	innerOverflow := []string{"0"}
	var sb strings.Builder
	for j, row := range rotated {
		// each row begins with starting with 0.
		carry := innerOverflow[0]
		if len(innerOverflow) > 1 {
			innerOverflow = innerOverflow[j:]
		} else {
			innerOverflow = []string{"0"}
		}

		// iterate over the numbers to get a grand total, and populate the overflow slice too.
		for _, col := range convertEmptyStringsToZeroes(row) {
			// do the thing infinitely until the sum is only one character long!
			carry = sum[carry][col]

			if len(carry) == 1 {
				continue
			}
			// if the length of carry is not one, it is two because the sum of two single digit numbers can only
			// result in a two digit number.
			// save the innerOverflow value that comes from the carry.
			innerOverflow[0] = sum[string(carry[0])][innerOverflow[0]]

			innerOverflow = carryInnerOverflow(innerOverflow)
			carry = string(carry[1])
		}
		sb.WriteString(carry)
	}

	for _, i := range innerOverflow {
		sb.WriteString(i)
	}

	result := strings.TrimLeft(reverseString(sb.String()), "0")
	if result == "" {
		result = "0"
	}

	return result
}

func convertEmptyStringsToZeroes(nums []string) []string {
	for i, num := range nums {
		if num == "" {
			nums[i] = "0"
		}
	}
	return nums
}

func rotateMatrix(input [][]string) [][]string {
	// we turn an 8 wide 2 deep matrix into a 2 wide 8 deep one.
	width := len(input[0])
	height := len(input)
	rotated := make([][]string, width)
	for i := range rotated {
		rotated[i] = make([]string, height)
	}

	for h, row := range input {
		for w, char := range row {
			rotated[w][h] = char
		}
	}
	return rotated
}

func carryInnerOverflow(innerOverflow []string) []string {
	// propagate the innerOverflows internally.
	lio := len(innerOverflow) - 1
	ouf := make([]string, len(innerOverflow))
	overflow := "0"
	for i, number := range innerOverflow {
		iofhelper := "0"
		if len(number) != 1 {
			iofhelper = string(number[0])
			number = string(number[1])
		}

		// this comes from the previous column.
		p1 := sum[overflow][number]

		// if the string of the sum of the carried overflow and the current number in the column is one digit long,
		// break and carry on.
		if len(p1) == 1 {
			ouf[i] = p1
		} else {
			ouf[i] = string(p1[1])
			iofhelper = sum[iofhelper][string(p1[0])]
		}

		// if the innerOverflow slice is not big enough, add a new element to the end
		if i+1 > lio && iofhelper != "0" {
			ouf = append(ouf, iofhelper)
			break
		}

		overflow = iofhelper
	}
	return ouf
}
