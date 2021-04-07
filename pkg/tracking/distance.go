package tracking

type SourceRelativeDistance struct {
	Sat []Satellite `json:"satellites"`
}

type Satellite struct {
	ID       string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}
