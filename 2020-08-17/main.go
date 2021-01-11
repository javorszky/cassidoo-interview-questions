package aug172020

import (
	"fmt"
	"math/rand"
)

func Task() {
	input := []int{1, 2, 0, 1, 0, 0, 3, 6}
	output := moveZeroes(input)

	fmt.Printf("17th August 2020\n---------------\n"+
		"original: %v\n"+
		"moved:    %v\n\n", input, output)

	//lenghts := []uint32{10, 100, 1000, 10000, 100000}
	//for _, v := range lenghts {
	//	fmt.Printf("%#v\n\n", generateSlice(v))
	//}
}

// moveZeroes is the solution to the interview question. Takes a slice, returns a slice with zeroes at the end.
func moveZeroes(nums []int) []int {
	moved := make([]int, 0, len(nums))
	for _, v := range nums {
		if v == 0 {
			continue
		}
		moved = append(moved, v)
	}
	moved = append(moved, make([]int, len(nums)-len(moved), len(nums)-len(moved))...)

	return moved
}

// generateSlice is a utility func used to generate the slices in the generated.go file.
func generateSlice(length uint32) []int {
	nums := make([]int, 0, length)
	for i := uint32(0); i < length; i++ {
		nums = append(nums, rand.Intn(36))
	}
	return nums
}
