package memory

import (
	"github.com/soulmarus/quasar-operation/pkg/tracking"
)

// Memory storage keeps data in memory
type Storage struct {
	stations []Station
	srds     map[string]Satellite
}

// Add saves the given station to the repository
func (m *Storage) AddStation(b tracking.Station) error {
	newS := Station{
		ID:  b.ID,
		Pos: Position{b.Pos.X, b.Pos.Y},
	}
	m.stations = append(m.stations, newS)

	return nil
}

// Get returns a station with the specified ID
func (m *Storage) GetStation(id string) (tracking.Station, error) {
	var station tracking.Station

	for i := range m.stations {

		if m.stations[i].ID == id {
			station.ID = m.stations[i].ID
			station.Pos.X = m.stations[i].Pos.X
			station.Pos.Y = m.stations[i].Pos.Y

			return station, nil
		}
	}

	return station, tracking.ErrNotFound
}

// Add saves or updates if existing the given source relative distance to a satellite to the repository
func (m *Storage) AddSourceRelativeDistance(b tracking.Satellite) error {
	newS := Satellite{
		ID:       b.ID,
		Distance: b.Distance,
		Message:  b.Message,
	}
	m.srds[newS.ID] = newS

	return nil
}

// Get returns a source relative distance to a satellite with the specified ID
func (m *Storage) GetSourceRelativeDistance(id string) (tracking.Satellite, error) {
	var srd tracking.Satellite

	if val, ok := m.srds[id]; ok {
		srd.ID = val.ID
		srd.Distance = val.Distance
		srd.Message = val.Message

		return srd, nil
	}

	return srd, tracking.ErrNotFound
}

// Get returns all source relative distances currently stored
func (m *Storage) GetAllSourceRelativeDistance() ([]tracking.Satellite, error) {
	var srds []tracking.Satellite

	for _, val := range m.srds {
		srd := tracking.Satellite{
			ID:       val.ID,
			Distance: val.Distance,
			Message:  val.Message,
		}
		srds = append(srds, srd)
	}

	return srds, nil
}

// InitializeSampleSrds create the map structure
func (m *Storage) InitializeSampleSrds() {
	m.srds = make(map[string]Satellite)
}
