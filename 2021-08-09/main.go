package august082021

import (
	"fmt"
)

func Tasks() {
	gridSize := 5
	mines := []Coordinate{
		{1, 3},
		{3, 5},
		{2, 4},
	}
	fmt.Printf("A minesweeper grid with size %d and mines at %v will look like this\n\n", gridSize, mines)
	fmt.Println(generateMinesweeper(gridSize, mines))
}

// generateMinesweeper takes a grid size integer, generates a 1-based square field, places the mines on the field, and
// annotates the adjacent fields by how many mines there are next to them, and returns them as a single string.
func generateMinesweeper(gridSize int, mines []Coordinate) string {
	f, err := NewField(gridSize, mines)
	if err != nil {
		return fmt.Sprintf("an error occurred: %s", err)
	}
	return f.String()
}
