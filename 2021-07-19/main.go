package july192021

import "fmt"

func Tasks() {
	fmt.Printf("sum of subs:\n%#v\n", subArraySum([]int{10, 2, -2, -20, 10}, -10))
}

func subArraySum(arr []int, n int) [][]int {
	subs := make([][]int, 0)
	for i := range arr {
		sub := arr[i:]
		for j := 0; j <= len(sub); j++ {
			s := addNumbers(sub[:j])
			if s == n {
				subs = append(subs, sub[:j])
			}
		}
	}
	return subs
}

// addNumbers is a utility function with the sole purpose of summing an array of ints.
func addNumbers(arr []int) int {
	acc := 0
	for _, i := range arr {
		acc = acc + i
	}
	return acc
}
