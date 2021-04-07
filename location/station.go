package location

type Station struct {
	x, y float64
	name string
}

var Stations = map[string]Station{
	"kenobi": {
		-500.0, -200.0, "kenobi",
	},
	"skywalker": {
		100.0, -100.0, "skywalker",
	},
	"sato": {
		500.0, 100.0, "sato",
	},
}
