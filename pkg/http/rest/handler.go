package rest

import (
	"encoding/json"
	"net/http"

	"github.com/soulmarus/quasar-operation/pkg/tracking"

	"github.com/julienschmidt/httprouter"
)

func Handler(t tracking.Service) http.Handler {
	router := httprouter.New()

	router.POST("/topsecret", getStationInfo(t))
	router.POST("/topsecret_split/:id", addSourceRelativeDistance(t))
	router.GET("/topsecret_split", getStationInfoSplit(t))

	return router
}

// getStationInfo returns a handler for POST /topsecret requests
func getStationInfo(t tracking.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var sourceDistances tracking.SourceRelativeDistance
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&sourceDistances); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sourceInfo, err := t.TrackSource(sourceDistances)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sourceInfo)
	}
}

// addSourceRelativeDistance returns a handler for POST /topsecret_split:id requests
func addSourceRelativeDistance(t tracking.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var newSatellite tracking.Satellite
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&newSatellite)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newSatellite.ID = p.ByName("id")

		t.AddSourceRelativeDistance(newSatellite)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New source distance recorded")
	}
}

// getStationInfoSplit returns a handler for GET /topsecret_split requests
func getStationInfoSplit(t tracking.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		sourceInfo, err := t.TrackSplitSource()

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sourceInfo)
	}
}
