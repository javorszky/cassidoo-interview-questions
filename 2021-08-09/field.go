package august082021

import (
	"fmt"
	"strings"
)

const (
	mine FieldType = iota
	zero
	one
	two
	three
	four
	five
	six
	seven
	eight
)

// Field is a representation of playing field.
type Field struct {
	size   int
	fields map[Coordinate]FieldType
}

func (f Field) addMine(coordinate Coordinate) (Field, error) {
	fieldValue, ok := f.fields[coordinate]
	if !ok {
		return f, fmt.Errorf("coordinate {%d, %d} does not exist on field with grid size %d", coordinate[0], coordinate[1], f.size)
	}

	if fieldValue == mine {
		return f, fmt.Errorf("trying to add a mine to a field that's already a mine at coordinate {%d, %d}", coordinate[0], coordinate[1])
	}

	f.fields[coordinate] = mine

	f, err := f.incrementAdjacents(coordinate)
	if err != nil {
		return f, fmt.Errorf("trying to increment adjacents while adding a mine for coordinate {%d, %d}: %w", coordinate[0], coordinate[1], err)
	}

	return f, nil
}

func (f Field) incrementAdjacents(coordinate Coordinate) (Field, error) {
	// i is the horizontal movement, minus 1, zero, or plus 1 of coordinate[0].
	for i := -1; i < 2; i++ {
		// j is the vertical movement, minus 1, zero, or plus 1 of coordinate[1].
		for j := -1; j < 2; j++ {
			// skip if it's the coordinate itself.
			if i == 0 && j == 0 {
				continue
			}

			xNew := coordinate[0] + i
			yNew := coordinate[1] + j

			// the adjacent field in question is out of bounds, ignore.
			if xNew < 1 || yNew < 1 || xNew > f.size || yNew > f.size {
				continue
			}

			cNew := Coordinate{xNew, yNew}

			fieldValue, ok := f.fields[cNew]

			// adjacent field does not exist, even though it's not out of bounds. ???!!?!
			if !ok {
				return f, fmt.Errorf("the adjacent field with coordinates {%d, %d} did not exist. It should have", xNew, yNew)
			}

			// adjacent field is a mine, ignore.
			if fieldValue == mine {
				continue
			}

			f.fields[cNew] = fieldValue + 1
		}
	}

	return f, nil
}

// String prints the field into a string grid
func (f Field) String() string {
	var sb strings.Builder

	for i := 1; i < f.size+1; i++ {
		for j := 1; j < f.size+1; j++ {
			sb.WriteString(f.fields[Coordinate{j, i}].String())
		}
		sb.WriteString("\n")
	}
	return strings.TrimRight(sb.String(), "\n")
}

// NewField takes the input data and generates a Field, or returns an error if it couldn't generate the field due to
// either empty list of mines or a mine being out of bounds.
func NewField(gridSize int, fields []Coordinate) (Field, error) {
	fieldMap := make(map[Coordinate]FieldType, gridSize*gridSize)
	for i := 1; i < 1+gridSize; i++ {
		for j := 1; j < 1+gridSize; j++ {
			fieldMap[Coordinate{i, j}] = zero
		}
	}

	f := Field{
		size:   gridSize,
		fields: fieldMap,
	}

	for _, c := range fields {
		t, err := f.addMine(c)
		if err != nil {
			return Field{}, fmt.Errorf("problem while adding mines to the new empty field: %w", err)
		}

		f = t
	}

	return f, nil
}
