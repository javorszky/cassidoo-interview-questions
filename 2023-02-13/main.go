package feb132023

import (
	"fmt"
	"log"
	"strings"
)

func Task() {
	for i := 0; i < 101; i++ {
		switch {
		case i > 99:
			// this is just the number 100
			oneSlice := getNumber(1)
			zeroSlice := getNumber(0)
			for i := range oneSlice {
				if i == 0 {
					fmt.Println("")
					continue
				}

				fmt.Printf("%s\t%s\t%s\n", oneSlice[i], zeroSlice[i], zeroSlice[i])
			}
		case i > 9:
			ten := i / 10
			one := i - (10 * ten)
			tSlice := getNumber(ten)
			oneSlice := getNumber(one)
			for i := range tSlice {
				if i == 0 {
					fmt.Println("")
					continue
				}

				fmt.Printf("%s\t%s\n", tSlice[i], oneSlice[i])
			}
		default:
			sloice := getNumber(i)
			for _, l := range sloice {
				fmt.Println(l)
			}
		}
	}
}

func getNumber(n int) []string {
	number, ok := numberMapping[n]
	if !ok {
		log.Fatalf("I can only deal with digits 0-9, you gave me %d", n)
	}

	return strings.Split(number, "\n")
}
