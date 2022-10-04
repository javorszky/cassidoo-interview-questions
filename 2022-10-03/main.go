package oct032022

import (
	"errors"
	"fmt"
)

func Tasks() {
	var a, b, n uint64

	a, b, n = 10, 20, 5
	got, err := fibLike(a, b, n)
	printThis(a, b, n, got, err)

	a, b = 3, 7
	got, err = fibLike(a, b, n)
	printThis(a, b, n, got, err)
}

func printThis(a, b, n uint64, s []uint64, err error) {
	if err != nil {
		fmt.Printf("could not calculate fibonacci-like sequence for %d, %d with %d length because %s\n", a, b, n, err.Error())
		return
	}

	fmt.Printf("inputs %d, %d, %d yielded the following sequence\n%v\n", a, b, n, s)

}

func fibLike(a, b, n uint64) ([]uint64, error) {
	if n < 2 {
		return nil, errors.New("sequence too short")
	}

	seq := make([]uint64, n, n)
	seq[0] = a
	seq[1] = b

	for i := uint64(2); i < n; i++ {
		a, b = b, a+b
		if b < a {
			return nil, errors.New("wrapped around, uint overflow")
		}

		seq[i] = b
	}

	return seq, nil
}
