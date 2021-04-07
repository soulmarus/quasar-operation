package tracking

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	EPSILON = 1
)

// ErrNotFound is used when some resource is not found
var ErrNotFound = errors.New("unable to locate the source of the messages")

// ErrSatNotFound is used when the sat ID informed doesn't exist.
var ErrSatNotFound = errors.New("unable to locate satelite")

// Service provides station tracking operations.
type Service interface {
	TrackSource(SourceRelativeDistance) (SourceInfo, error)
	AddSampleStations([]Station)
	AddSourceRelativeDistance(b Satellite) error
	TrackSplitSource() (SourceInfo, error)
	InitializeSampleSrds()
}

// Repository provides access to tracking repository.
type Repository interface {
	// AddStation saves a given tracking to the repository
	AddStation(Station) error
	GetStation(id string) (Station, error)
	AddSourceRelativeDistance(Satellite) error
	GetSourceRelativeDistance(id string) (Satellite, error)
	GetAllSourceRelativeDistance() ([]Satellite, error)
	InitializeSampleSrds()
}

type service struct {
	r Repository
}

// NewService creates an tracking service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddSampleStations adds some sample Stations to the database
func (s *service) AddSampleStations(st []Station) {
	for _, sst := range st {
		_ = s.r.AddStation(sst)
	}
}

// TrackSource given a list of distances calculate the position and determine the real message expecting some lags
func (s *service) TrackSource(srd SourceRelativeDistance) (SourceInfo, error) {
	return trackSource(srd, s)
}

// TrackSource given a list of distances calculate the position and determine the real message expecting some lags
func trackSource(srd SourceRelativeDistance, s *service) (SourceInfo, error) {
	// validate if there is 3 valid sources
	err := validate(srd)

	if err != nil {
		return SourceInfo{}, err
	}

	sat1 := srd.Sat[0]
	sat2 := srd.Sat[1]
	sat3 := srd.Sat[2]

	station1, err := s.r.GetStation(sat1.ID)

	if err != nil {
		return SourceInfo{}, fmt.Errorf("station %s %v", sat1.ID, err)
	}

	station2, err := s.r.GetStation(sat2.ID)

	if err != nil {
		return SourceInfo{}, fmt.Errorf("station %s %v", sat2.ID, err)
	}

	station3, err := s.r.GetStation(sat3.ID)

	if err != nil {
		return SourceInfo{}, fmt.Errorf("station %s %v", sat3.ID, err)
	}

	x, y, err := calculateThreeCircleIntersection(station1.Pos.X, station1.Pos.Y, sat1.Distance,
		station2.Pos.X, station2.Pos.Y, sat2.Distance,
		station3.Pos.X, station3.Pos.Y, sat3.Distance)

	if err != nil {
		return SourceInfo{}, err
	}

	return SourceInfo{Pos: Position{x, y}, Message: getMessage(sat1.Message, sat2.Message, sat3.Message)}, nil
}

func validate(srd SourceRelativeDistance) error {
	if len(srd.Sat) < 3 {
		return fmt.Errorf("at least three distance points are required to triangulate the position")
	}

	seen := make(map[string]bool)

	for _, sat := range srd.Sat {
		if seen[sat.ID] {
			return fmt.Errorf("there are duplicated satelites ID")
		}

		seen[sat.ID] = true
	}

	return nil
}

// calculateThreeCircleIntersection calculate the position on a 2D plane from three points and the radius (distance) relative to the source of the signal
// It returns the data points if it's possible to calculate otherwise it will return any error encountered
func calculateThreeCircleIntersection(x0 float64, y0 float64, r0 float64,
	x1 float64, y1 float64, r1 float64,
	x2 float64, y2 float64, r2 float64) (float64, float64, error) {
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
		fmt.Printf("impossible to triangulate position circles do not intersect \n")
		return 0.0, 0.0, fmt.Errorf("impossible to triangulate position circles do not intersect")
	}

	if d < math.Abs(r0-r1) {
		/* no solution. one circle is contained in the other */
		fmt.Printf("impossible to triangulate position one circle is contained in the other \n")
		return 0.0, 0.0, fmt.Errorf("impossible to triangulate position one circle is contained in the other")
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
		fmt.Printf("INTERSECTION Circle1 AND Circle2 AND Circle3:: (%f,%f) \n", intersectionPoint1_x, intersectionPoint1_y)
		return intersectionPoint1_x, intersectionPoint1_y, nil
	} else if math.Abs(d2-r2) < EPSILON {
		fmt.Printf("INTERSECTION Circle1 AND Circle2 AND Circle3:: (%f,%f) \n", intersectionPoint2_x, intersectionPoint2_y)
		return intersectionPoint2_x, intersectionPoint2_y, nil
	} else {
		fmt.Println("INTERSECTION Circle1 AND Circle2 AND Circle3:", "NONE")
		return 0.0, 0.0, fmt.Errorf("impossible to triangulate position can't intersect the three points")
	}
}

func getMessage(messages ...[]string) (msg string) {
	finalMessage := ""
	seen := make(map[string]bool)
	maxLen := determineMaxLength(messages...)

	for i := 0; i < maxLen; i++ {
		for _, message := range messages {
			if i < len(message) {
				if message[i] != "" {
					if seen[message[i]] {
						continue
					}
					finalMessage += message[i] + " "
					seen[message[i]] = true
				}
			}
		}
	}
	return strings.TrimSpace(finalMessage)
}

func determineMaxLength(messages ...[]string) int {
	maxLen := 0
	for _, message := range messages {
		len := len(message)
		if maxLen < len {
			maxLen = len
		}
	}
	return maxLen
}

// AddSourceRelativeDistance persists the given beer(s) to storage
func (s *service) AddSourceRelativeDistance(b Satellite) error {

	err := s.r.AddSourceRelativeDistance(b)

	if err != nil {
		return nil
	}

	return nil
}

// TrackSource given a list of distances that came in diferrent
// time calculate the position and determine the real message expecting some lags
func (s *service) TrackSplitSource() (SourceInfo, error) {
	srds, err := s.r.GetAllSourceRelativeDistance()

	if err != nil {
		return SourceInfo{}, err
	}

	srd := SourceRelativeDistance{
		Sat: srds,
	}

	return trackSource(srd, s)
}

// InitializeSampleSrds create the map structure
func (s *service) InitializeSampleSrds() {
	s.r.InitializeSampleSrds()
}
