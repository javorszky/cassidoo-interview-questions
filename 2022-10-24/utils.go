package oct242022

type side int

const (
	noSide side = iota
	leftSide
	rightSide
)

func splitSlice(s []uint, lean side) ([]uint, uint, []uint) {
	if lean == noSide {
		lean = leftSide
	}

	l := len(s)
	if l == 0 {
		return []uint{}, 0, []uint{}
	}

	// if it's odd, return the correct split
	if l%2 == 1 {
		return s[:l/2], s[l/2], s[l/2+1:]
	}

	half := l / 2

	if lean == leftSide {
		return s[:half], s[half], s[half+1:]
	}

	return s[:half-1], s[half-1], s[half:]
}

func build(s []uint, lean side) *node {
	if len(s) == 0 {
		return nil
	}

	left, val, right := splitSlice(s, lean)
	return &node{
		value: val,
		left:  build(left, leftSide),
		right: build(right, rightSide),
	}
}
