package sep212020

import (
	"fmt"
	"math"
)

func Task() {
	//input := []int{5, 8, -5}
	//input := []int{4, 0, -0}
	input := []int{2, 4, 8, -9, -8, -1, -4, 7, -5, 5, 0, 1, -3, -8, 7, 14, -8, -22, -7}
	endRoids := asteroids(input)
	fmt.Printf("21st September 2020\n-------------------\n"+
		"Beginning asteroids:  %v\n"+
		"Final state of roids: %v\n\n", input, endRoids)
}

func asteroids(roids []int) []int {
	newRoids := make([]int, 0)
	//fmt.Printf("roids to collide: %v\n", roids)
	for index, value := range roids {
		//fmt.Printf("---------------\nstarting a new iteration. Current new roids:\n%v\n\n", newRoids)
		if index == 0 {
			newRoids = append(newRoids, value)
			//fmt.Printf("index was zero, new roids: %v\n", newRoids)
			// skip the first item, we don't have a "previous" to compare with.
			continue
		}

		previous := newRoids[len(newRoids)-1]

		//fmt.Printf("checking this %d and previous %d\n", value, previous)

		// are the signs different?
		if value*previous > 0 {
			// they are the same direction, won't collide, skip.
			newRoids = append(newRoids, value)
			//fmt.Printf("their signs were the same, new roids: %v\n", newRoids)
			continue
		}

		if value+previous == 0 {
			// Both values are the same, delete last one, do not add the current one.
			newRoids = newRoids[:len(newRoids)-1]
			//fmt.Printf("both this %d and previous %d are the same size. They both get destroyed. New roids: \n%v\n\n", value, previous, newRoids)
		} else if math.Abs(float64(value)) > math.Abs(float64(previous)) {
			// This value is bigger than the previous one. Remove the previous one, add current one.
			newRoids = newRoids[:len(newRoids)-1]
			newRoids = append(newRoids, value)
			//fmt.Printf("this %d was bigger than previous %d, removing the last element of newroids, and adding this."+
			//	"\nnew roids therefore:\n%v\n\n", value, previous, newRoids)
			//fmt.Println("we messed with the existing list, so recursively dealing with new roids!")
			newRoids = asteroids(newRoids)
		} else {
			// Previous value is bigger, do not remove previous, do not add this one.
			//fmt.Printf("previous %d was bigger than this %d. Leaving new roids as is:\n%v\n\n", previous, value, newRoids)
		}
	}
	return newRoids
}
