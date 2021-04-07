package location

import (
	"fmt"
	"testing"
)

func TestCalculateThreeCircleIntersection(t *testing.T) {
	var tests = []struct {
		x0 float64
		y0 float64
		r0 float64
		x1 float64
		y1 float64
		r1 float64
		x2 float64
		y2 float64
		r2 float64
		w  bool
	}{{-500.0, -200.0, 424.264069, 100.0, -100.0, 360.555128, 500.0, 100.0, 700.0, true},
		{-500.0, -200.0, 100.0, 100.0, -100.0, 115.5, 500.0, 100.0, 142.7, true}}

	for _, test := range tests {
		fmt.Println("Calculate:", CalculateThreeCircleIntersection(test.x0, test.y0, test.r0, test.x1, test.y1, test.r1,
			test.x2, test.y2, test.r2))
	}
}
