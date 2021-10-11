package october112021

import (
	"fmt"
	"strings"
)

func Tasks() {
	for _, v := range []int{14, 5, 98, 13, 232, 74} {
		testAndPrint(v)
	}
}

func testAndPrint(n int) {
	fmt.Printf("The oddity of the number %d is %t\n", n, isOdious(n))
}

func isOdious(n int) bool {
	return strings.Count(fmt.Sprintf("%b", n), "1")%2 == 1
}
