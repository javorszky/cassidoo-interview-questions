package jun142021

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const star = "*\n"
const space = " "

type padder func(int) string

func Task() {
	fmt.Printf("drawing arrow left 5:\n%s\n", drawArrow("left", 5))
	fmt.Printf("drawing arrow right 8:\n%s\n", drawArrow("right", 8))
}

func drawArrow(direction string, depth int) string {
	var p padder
	switch direction {
	case "left":
		p = leftPad(depth)
	case "right":
		p = rightPad(depth)
	default:
		panic(fmt.Errorf("invalid direction. Expected 'left' or 'right', got '%s'", direction))
	}

	var sb strings.Builder

	drawArrowRecursive(depth, p, &sb)

	return strings.TrimRight(sb.String(), "\n")
}

func drawArrowRecursive(depth int, p padder, sb io.Writer) {
	padding := p(depth)

	if depth-1 == 0 {
		_, err := sb.Write([]byte(padding + star))
		if err != nil {
			panic(errors.New("could not write to strings builder at apex of arrow"))
		}
		return
	}

	_, err := sb.Write([]byte(padding + star))
	if err != nil {
		panic(fmt.Errorf("could not write to strings builder on the top half of arrow at depth %d", depth))
	}

	drawArrowRecursive(depth-1, p, sb)

	_, err = sb.Write([]byte(padding + star))
	if err != nil {
		panic(fmt.Errorf("could not write to strings builder on the bottom half of arrow at depth %d", depth))
	}
}

func leftPad(_ int) func(int) string {
	return func(depth int) string {
		if depth-1 < 0 {
			return ""
		}
		return strings.Repeat(space, depth-1)
	}
}

func rightPad(maxDepth int) func(int) string {
	return func(depth int) string {
		if depth-1 < 0 {
			return strings.Repeat(space, maxDepth)
		}
		return strings.Repeat(space, maxDepth-depth)
	}
}
