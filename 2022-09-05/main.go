package sept052022

import (
	"errors"
	"fmt"
)

func Tasks() {
	eh := fromTo(5, 7)

	for i := 0; i < 4; i++ {
		c, err := eh()
		fmt.Printf("call %d: returned %d, error? %v\n", i, c, err)
	}
}

type genFunc func() (int, error)

func fromTo(a, b int) genFunc {
	called := 0
	return func() (int, error) {
		if a+called > b {
			return 0, errors.New("called too many times")
		}

		r := a + called
		called += 1

		return r, nil
	}
}
