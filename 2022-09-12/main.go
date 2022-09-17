package sept122022

import (
	"fmt"
	"log"
)

func Tasks() {
	logoDimensions := dimension{
		w: 5,
		h: 5,
	}
	logoStartCoord := coordinate{
		x: 1,
		y: 2,
	}
	screenSize := dimension{
		w: 86,
		h: 100,
	}

	ok, howMany, err := cornerHit(logoDimensions, logoStartCoord, screenSize)
	if err != nil {
		log.Fatalf("some error occurred here: %s", err.Error())
	}
	if ok {
		fmt.Printf("A logo with the size of [%d, %d] and top left corner start position of [%d, %d] will hit a corner in %d hits in a screen of [%d, %d]\n", logoDimensions.w, logoDimensions.h, logoStartCoord.x, logoStartCoord.y, howMany, screenSize.w, screenSize.h)
	} else {
		fmt.Printf("A logo with the size of [%d, %d] and top left corner start position of [%d, %d] will never hit a corner in a screen of [%d, %d]\n", logoDimensions.w, logoDimensions.h, logoStartCoord.x, logoStartCoord.y, screenSize.w, screenSize.h)
	}
}

// dimension describes a thing with a size. w is for horizontal size (width), h for vertical (height).
type dimension struct {
	w, h int
}

// canContain on a dimension takes another dimension, and returns a bool if the other dimension can fit inside the first
// one.
func (d dimension) canContain(other dimension) bool {
	return d.w-other.w > 0 && d.h-other.h > 0
}

// coordinate describes a point in space. x is the horizontal position, y is the vertical position.
type coordinate struct {
	x, y int
}

// horizontal <-- x -->
// vertical ^
//          |  calculate what x value the path it started with based on the locoCoord.
//          y
//          |  For this we're using the top left corner, and we're reducing the screen
//          v  size by the logo size, so if screen size is 800 x 600, and the logo is
//             100 x 70, then the virtual screen the top left corner can take is only
//             700 x 530.
//
// Let's also make an infinite grid of 700 x 530, and then calculate where the x coord
// will be where y = 0.
func cornerHit(logoSize dimension, logoCoord coordinate, screenSize dimension) (bool, int, error) {
	if !screenSize.canContain(logoSize) {
		return false, 0, fmt.Errorf("logo (%dx%d) does not fit in screen (%dx%d)",
			logoSize.w, logoSize.h,
			screenSize.w, screenSize.h,
		)
	}
	fmt.Println("screensize can contain logo")

	fmt.Printf("Logo size: [%d, %d]\n", logoSize.w, logoSize.h)
	fmt.Printf("screen size: [%d, %d]\n", screenSize.w, screenSize.h)

	virtualScreenSize := dimension{
		w: screenSize.w - logoSize.w,
		h: screenSize.h - logoSize.h,
	}

	fmt.Printf("virtual screen size is [%d, %d]\n", virtualScreenSize.w, virtualScreenSize.h)

	xCoord := findXCoordinateForY0(virtualScreenSize, logoCoord)

	fmt.Printf("starting x coordinate is [%d, 0] from start position [%d, %d]\n\n", xCoord, logoCoord.x, logoCoord.y)

	// now comes the fun part. Does 0 + virtualscreen width and xCoord + logo width have overlapping points?
	// start with the current x coordinate on the top edge
	currentX := xCoord

	// keep track of how far x needs to move
	cumulativeX := xCoord

	// keep track of x coordinates seen on the top edge. The moment we land on a point where we've already been, we'll
	// begin a new loop, which means there's no point in calculating further.
	seenX := make(map[int]struct{})

	loop := 0

	// start trying
	for {
		loop++
		// next step in x coordinate
		currentX = (currentX + virtualScreenSize.h) % virtualScreenSize.w

		// keep track how far to the side we're moving, important for bump calculations.
		cumulativeX += virtualScreenSize.h

		fmt.Printf("Loop %d: currentx: %d, cumx: %d\n", loop, currentX, cumulativeX)

		if currentX == 0 {
			fmt.Printf("Loop %d: we have hit a corner!\n", loop)
			// we have hit a corner!
			return true, calculateBumps(logoCoord, virtualScreenSize, cumulativeX), nil
		}

		if _, ok := seenX[currentX]; ok {
			fmt.Printf("Loop %d: we have seen %d previously, is an infinite loop, exiting\n", loop, currentX)
			// we have started to loop and have not reached a 0 by this time, which means we will never ever hit a
			// corner
			return false, 0, nil
		}

		// we haven't seen this yet, but also haven't seen the corner, so store the x, and try again.
		seenX[currentX] = struct{}{}
		fmt.Printf("Loop %d: storing currentx %d to seen map\n", loop, currentX)
	}
}

// findXCoordinateForY0 returns the x (horizontal) coordinate if the point is on the top edge.
//
// To do that, because the direction is always a 1x1 in any of the 4 directions, we have to take away the y coordinate
// from the x coordinate, and then modulo adjust in case we get a negative number. Here we assume that the direction is
// going to be to bottom right.
//
// We get an unadjusted X coordinate, which may be super negative in case the screen is 230 wide and 7000 tall. We need
// to bring that back into the one screen given infinite tiling.
//
// To do that, we take the modulo of unadjusted x / screen width, then add the width to the result, and then modulo that
// result by the width again. This way both negative and positive results will work. See examples:
//
// --- negative ---
// -5434 % 230 = -144
//  -144 + 230 =   86
//    86 % 230 =   86
//
// --- positive ---
//     7 % 230 =    7
//     7 + 230 =  237
//   237 % 230 =    7
func findXCoordinateForY0(screenSize dimension, logoCoord coordinate) int {
	xUnadjusted := logoCoord.x - logoCoord.y
	xAdjusted := (xUnadjusted%screenSize.w + screenSize.w) % screenSize.w
	return xAdjusted
}

// calculateBumps will tell us how many bumps it took to arrive to a corner piece. This is only called if we've actually
// hit a corner piece.
//
// First thing is to calculate how many bumps we've gone backwards to arrive from the starter position to the point on
// the top edge using findXCoordinateForY0.
//
// Then it's a calculation of how many times we've crossed both the side on the right, and the bottom one.
func calculateBumps(startPosition coordinate, virtualScreenSize dimension, cumulativeX int) int {
	// first, let's calculate how many bumps we've gone backwards. By definition, we've not crossed any of the bottom
	// edge.
	initialBumps := ((virtualScreenSize.w - startPosition.x) + startPosition.y) / virtualScreenSize.w
	fmt.Printf("BUMPFINDER: initial bumps is %d\n", initialBumps)

	// then figure out how many times we bumped into (crossed) the bottom edge
	yBumps := cumulativeX / virtualScreenSize.h
	fmt.Printf("BUMPFINDER: yBumps bumps is %d\n", yBumps)

	// and how many times we bumped into (crossed) the right edge
	xBumps := cumulativeX / virtualScreenSize.w
	fmt.Printf("BUMPFINDER: xBumps is %d\n", xBumps)

	return xBumps + yBumps - initialBumps
}
