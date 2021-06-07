package jun072021

import "fmt"

func Task() {
	//fmt.Printf("the sum of numbers 1,2,3 is %d", lonelyNumber(1, 2, 3))
	fmt.Printf("the sum of numbers 6,6,4 is %d", lonelyNumber(6, 6, 4))
	//fmt.Printf("the sum of numbers 7,7,7 is %d", lonelyNumber(7, 7, 7))
}

type ints map[int]int

func lonelyNumber(a, b, c int) int {
	i := make(ints)
	i[a] = i[a] + 1
	i[b] = i[b] + 1
	i[c] = i[c] + 1

	product := 1

	for k, v := range i {
		if v > 1 {
			continue
		}
		product = product * k
	}
	return product
}
