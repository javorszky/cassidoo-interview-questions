package october182021

import (
	"errors"
	"fmt"
)

func Tasks() {
	arr := []string{"C", "D", "E", "F", "G", "H"}
	indices := []int{3, 0, 4, 1, 2, 5}
	ordered, err := reOrder(arr, indices)
	fmt.Printf("inputs:\n%v\n%v\n\n", arr, indices)
	fmt.Printf("ordered list is\n%#v\n\nand error is: %v\n", ordered, err)
}

func reOrder(arr []string, indices []int) (ordered []string, ordErr error) {
	orderedMap := make(map[int]string)

	defer func() {
		if r := recover(); r != nil {
			ordered = nil
		}
	}()

	if len(arr) != len(indices) {
		return nil, errors.New("character list and index list is not the same. They should be")
	}

	for i, j := range indices {
		orderedMap[j] = arr[i]
	}

	if len(orderedMap) != len(arr) {
		return nil, errors.New("resulting ordered list is not the same length as incoming character list. Index list had a duplicate in it")
	}

	ordered = make([]string, len(arr))

	for k, v := range orderedMap {
		ordered[k] = v
	}

	return ordered, nil
}

//
//
//> Given an array of objects A, and an array of indexes B, reorder the objects in array A with the given indexes in array B.
//>
//> Example:
//> ```shell
//> let a = [C, D, E, F, G, H];
//> let b = [3, 0, 4, 1, 2, 5];
//>
//> $ reorder(a, b) // a is now [D, F, G, C, E, H]
//> ```
