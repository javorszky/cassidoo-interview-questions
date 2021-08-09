package august082021

import "fmt"

// Coordinate is a 2 value array representing an x and y coordinate of a mine or field.
type Coordinate [2]int

// String renders the string representation of the 2 member array.
func (c Coordinate) String() string {
	return fmt.Sprintf("%d-%d", c[0], c[1])
}
