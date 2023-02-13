package oct242022

import (
	"fmt"
)

func Tasks() {
	printTree(8)
}

func printTree(n uint) []byte {
	fmt.Printf("making a slice that's %d long\n", n)
	s := make([]uint, n)
	for i := uint(0); i < n; i++ {
		fmt.Printf("-- iter %d with value %d\n", i, i)
		s[i] = i
	}

	enn := build(s, leftSide)

	fmt.Printf("%#v\n", enn)

	fmt.Println(enn)

	return nil
}
