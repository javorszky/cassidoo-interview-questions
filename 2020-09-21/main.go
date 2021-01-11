package sep212020

import (
	"fmt"
	"math"
)

func Task() {
	input := []int{2, 4, 8, -9, -8, -1, -4, 7, -5, 5, 0, 1, -3, -8, 7, 14, -8, -22, -7}
	endRoids := asteroids(input)
	fmt.Printf("21st September 2020\n-------------------\n"+
		"Beginning asteroids:  %v\n"+
		"Final state of roids: %v\n\n", input, endRoids)
}

func asteroids(roids []int) []int {
	newRoids := make([]int, 0)
	for index, value := range roids {
		if index == 0 {
			newRoids = append(newRoids, value)
			// skip the first item, we don't have a "previous" to compare with.
			continue
		}

		previous := newRoids[len(newRoids)-1]

		// are the signs different?
		if value*previous > 0 {
			// they are the same direction, won't collide, skip.
			newRoids = append(newRoids, value)
			continue
		}

		if value+previous == 0 {
			// Both values are the same, delete last one, do not add the current one.
			newRoids = newRoids[:len(newRoids)-1]
		} else if math.Abs(float64(value)) > math.Abs(float64(previous)) {
			// This value is bigger than the previous one. Remove the previous one, add current one.
			newRoids = newRoids[:len(newRoids)-1]
			newRoids = append(newRoids, value)
			newRoids = asteroids(newRoids)
		}
	}
	return newRoids
}
