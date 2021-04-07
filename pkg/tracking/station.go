package tracking

type SourceInfo struct {
	Pos     Position `json:"position"`
	Message string   `json:"message"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Station struct {
	ID  string   `json:"name"`
	Pos Position `json:"position"`
}
