package tracking

import (
	"testing"
)

func TestCalculateThreeCircleIntersection(t *testing.T) {
	var tests = []struct {
		x0   float64
		y0   float64
		r0   float64
		x1   float64
		y1   float64
		r1   float64
		x2   float64
		y2   float64
		r2   float64
		wx   float64
		wy   float64
		werr bool
	}{{-500.0, -200.0, 424.264069, 100.0, -100.0, 360.555128, 500.0, 100.0, 700.0, -200.0000001641386, 100.00000057153323, false},
		{-500.0, -200.0, 100.0, 100.0, -100.0, 115.5, 500.0, 100.0, 142.7, 0.0, 0.0, true}, // expect error because circles do not intersect
		{-500.0, -200.0, 100.0, 100.0, -100.0, 800.0, 500.0, 100.0, 142.7, 0.0, 0.0, true}, // expect error because there is a circle inside another circle
		{-500.0, -200.0, 350.0, 100.0, -100.0, 300.0, 500.0, 100.0, 600.0, 0.0, 0.0, true}, // expect error because intersecting the three points is not possible
	}

	for _, test := range tests {
		x, y, err := calculateThreeCircleIntersection(test.x0, test.y0, test.r0, test.x1, test.y1, test.r1,
			test.x2, test.y2, test.r2)

		if err != nil && !test.werr {
			t.Errorf("Wanted no error(%v) got(true) %v", test.werr, err)
		}

		if x != test.wx && y != test.wy {
			t.Errorf("Wanted (X,Y) (%v, %v) got (%v, %v)", x, y, test.wx, test.wy)
		}
	}
}

func TestGetMessage(t *testing.T) {
	var tests = []struct {
		messages [][]string
		w        string
	}{
		{[][]string{{"", "este", "es", "un", "mensaje"},
			{"este", "", "un", "mensaje"},
			{"", "este", "es", "", ""}}, "este es un mensaje"},
		{[][]string{{"", "este", "es", "un", "mensaje"},
			{"este", "", "un", "mensaje"},
			{"", "este", "es", "", "secreto"}}, "este es un mensaje secreto"},
		{[][]string{{"", "es", "", "un", "mensaje"},
			{"este", "", "", "", ""},
			{"", "", "", "", "secreto"}}, "este es un mensaje secreto"},
		{[][]string{{"este", "", "", "mensaje", ""},
			{"", "es", "", "", "secreto"},
			{"este", "", "un", "", ""}}, "este es un mensaje secreto"},
		{[][]string{{"este", "mensaje"},
			{"es", ""},
			{"un"}}, "este es un mensaje"},
		{[][]string{}, ""},
	}

	for _, test := range tests {
		got := getMessage(test.messages...)

		if got != test.w {
			t.Errorf("Got \"%s\" wanted \"%s\" \n", got, test.w)
		}
	}
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		srd SourceRelativeDistance
		err bool
	}{
		{srd: SourceRelativeDistance{Sat: []Satellite{{"kenobi", 100.0, []string{""}}, {"kenobi", 100.0, []string{""}}, {"kenobi", 100.0, []string{""}}}}, err: true},
		{srd: SourceRelativeDistance{Sat: []Satellite{{"kenobi", 100.0, []string{""}}, {"kenobi", 100.0, []string{""}}}}, err: true},
		{srd: SourceRelativeDistance{Sat: []Satellite{{"kenobi", 100.0, []string{""}}, {"sato", 100.0, []string{""}}, {"skywalker", 100.0, []string{""}}}}, err: false},
	}

	for _, test := range tests {
		err := validate(test.srd)

		if err != nil && !test.err {
			t.Errorf("Wanted no error(%v) got(true) %v", test.err, err)
		}
	}
}
