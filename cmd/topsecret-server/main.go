package main

import (
	"fmt"
	"log"
	"net/http"
	"src/github.com/soulmarus/mercadolibre/pkg/http/rest"
	"src/github.com/soulmarus/mercadolibre/pkg/storage/memory"
	"src/github.com/soulmarus/mercadolibre/pkg/tracking"
)

// StorageType defines available storage types
type Type int

const (
	// Memory will store data in memory
	// With room for improvement and maybe use a DB so I can scale the MS
	Memory Type = iota
)

func main() {
	// setting up the storage
	storageType := Memory // I Could parse a flag here

	var tracker tracking.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)
		tracker = tracking.NewService(s)
		tracker.AddSampleStations(
			[]tracking.Station{{ID: "kenobi", Pos: tracking.Position{X: -500.0, Y: -200.0}},
				{ID: "skywalker", Pos: tracking.Position{X: 100.0, Y: -100.0}},
				{ID: "sato", Pos: tracking.Position{X: 500.0, Y: 100.0}}},
		)
		tracker.InitializeSampleSrds()
	}

	// set up the HTTP server
	router := rest.Handler(tracker)

	fmt.Println("Top Secret Server is starting now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router)) // The port here could also be a flag
}
