package location

import (
	"fmt"
	"math"
)

const (
	EPSILON = 1
)

func CalculateThreeCircleIntersection(x0 float64, y0 float64, r0 float64,
	x1 float64, y1 float64, r1 float64,
	x2 float64, y2 float64, r2 float64) bool {
	var a, dx, dy, d, h, rx, ry float64
	var point2_x, point2_y float64

	/* dx and dy are the vertical and horizontal distances between
	* the circle centers.
	 */
	dx = x1 - x0
	dy = y1 - y0

	/* Determine the straight-line distance between the centers. */
	d = math.Hypot(dy, dx)

	/* Check for solvability. */
	if d > (r0 + r1) {
		/* no solution. circles do not intersect. */
		fmt.Printf("no solution. circles do not intersect. \n")
		return false
	}

	if d < math.Abs(r0-r1) {
		/* no solution. one circle is contained in the other */
		fmt.Printf(" no solution. one circle is contained in the other. \n")
		return false
	}

	/* 'point 2' is the point where the line through the circle
	 * intersection points crosses the line between the circle
	 * centers.
	 */

	/* Determine the distance from point 0 to point 2. */
	a = ((r0 * r0) - (r1 * r1) + (d * d)) / (2.0 * d)

	/* Determine the coordinates of point 2. */
	point2_x = x0 + (dx * a / d)
	point2_y = y0 + (dy * a / d)

	/* Determine the distance from point 2 to either of the
	* intersection points.
	 */
	h = math.Sqrt((r0 * r0) - (a * a))

	/* Now determine the offsets of the intersection points from
	* point 2.
	 */
	rx = -dy * (h / d)
	ry = dx * (h / d)

	/* Determine the absolute intersection points. */
	intersectionPoint1_x := point2_x + rx
	intersectionPoint2_x := point2_x - rx
	intersectionPoint1_y := point2_y + ry
	intersectionPoint2_y := point2_y - ry

	fmt.Printf("INTERSECTION Circle1 AND Circle2: (%f,%f) AND (%f,%f) \n", intersectionPoint1_x, intersectionPoint1_y, intersectionPoint2_x, intersectionPoint2_y)

	/* Lets determine if circle 3 intersects at either of the above intersection points. */
	dx = intersectionPoint1_x - x2
	dy = intersectionPoint1_y - y2
	d1 := math.Sqrt((dy * dy) + (dx * dx))

	dx = intersectionPoint2_x - x2
	dy = intersectionPoint2_y - y2
	d2 := math.Sqrt((dy * dy) + (dx * dx))

	if math.Abs(d1-r2) < EPSILON {
		fmt.Printf("INTERSECTION D1-R2 Circle1 AND Circle2 AND Circle3:: (%f,%f) \n", intersectionPoint1_x, intersectionPoint1_y)
		// can't intersect the three points so we should return not found
	} else if math.Abs(d2-r2) < EPSILON {
		fmt.Printf("INTERSECTION D2-R2 Circle1 AND Circle2 AND Circle3:: (%f,%f) \n", intersectionPoint2_x, intersectionPoint2_y)
	} else {
		fmt.Println("INTERSECTION Circle1 AND Circle2 AND Circle3:", "NONE")
		// can't intersect the three points so we should return not found
	}
	return true
}
