package oct242022

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	leftBranch  = `/`
	rightBranch = `\`
	spacer      = ` `
)

type node struct {
	value uint
	left  *node
	right *node
}

func (n *node) String() string {
	_, s, err := n.stringLeaves(0)
	if err != nil {
		return err.Error()
	}
	return s
}

func (n *node) stringLeaves(depth int) (int, string, error) {
	fmt.Printf("%s -- NEW ------------\n\n", strings.Repeat(spacer, 2*depth))
	// If this is a leaf node, then return an empty string that's 0 wide. The level above would be responsible to
	// balance it with the length of the other side.
	if n.left == nil && n.right == nil {
		return 2, "  ", nil
	}

	ul, ur, ll, lr := "", "", "", ""

	if n.left != nil {
		w, s, err := n.left.stringLeaves(depth + 1)
		if err != nil {
			return 0, "", errors.Wrap(err, "left string leaves")
		}

		ul = strings.Repeat(spacer, w/2) + leftBranch + strings.Repeat(spacer, w-w/2-1)
		ll = s
	}

	if n.right != nil {
		w, s, err := n.right.stringLeaves(depth + 1)
		if err != nil {
			return 0, "", errors.Wrap(err, "right string leaves")
		}
		ur = strings.Repeat(spacer, w-w/2-1) + rightBranch + strings.Repeat(spacer, w/2)
		lr = s
	}

	// Split the string returned from the branches by the line break, giving us many lines, one per level.
	llLines := strings.Split(ll, "\n")
	lrLines := strings.Split(lr, "\n")

	// Check whether the line widths within the sides are consistent. If not, return an error
	w := len(llLines[0])
	for _, l := range llLines {
		if w != len(l) {
			return 0, "", fmt.Errorf("")
		}
	}

	// Check how many levels deep each side goes
	lenL := len(llLines)
	lenR := len(lrLines)
	lineWidth := len(llLines[0])
	if lineWidth < len(lrLines[0]) {
		lineWidth = len(lrLines[0])
	}

	if ul == "" {
		ul = strings.Repeat(spacer, lineWidth)
	}

	if ur == "" {
		ur = strings.Repeat(spacer, lineWidth)
	}

	// fmt.Printf("%s -- left\n"+
	// 	"%s %#v\n\n"+
	// 	"%s -- right\n"+
	// 	"%s %#v\n\n",
	// 	strings.Repeat(spacer, 2*depth),
	// 	strings.Repeat(spacer, 2*depth), llLines,
	// 	strings.Repeat(spacer, 2*depth),
	// 	strings.Repeat(spacer, 2*depth), lrLines,
	// )

	smaller := lenL
	bigger := lenR
	longer := "right"
	if lenR < smaller {
		smaller = lenR
		bigger = lenL
		longer = "left"
	}

	aggregated := make([]string, bigger+1)

	for i := 0; i < smaller; i++ {
		aggregated[i] = llLines[i] + lrLines[i]
	}

	fmt.Printf("linewidth: %d\n\n", lineWidth)

	switch longer {
	case "left":
		for i := smaller; i < bigger; i++ {
			aggregated[i] = llLines[i] + strings.Repeat(spacer, lineWidth)
		}
	case "right":
		for i := smaller; i < bigger; i++ {
			aggregated[i] = strings.Repeat(spacer, lineWidth) + lrLines[i]
		}
	}

	all := ul + ur + "\n" + strings.Join(aggregated, "\n")

	return 2 * lineWidth, all, nil
}
