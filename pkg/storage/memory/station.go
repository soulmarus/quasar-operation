package memory

type Station struct {
	ID  string
	Pos Position
}

type Position struct {
	X float64
	Y float64
}

type Satellite struct {
	ID       string
	Distance float64
	Message  []string
}
