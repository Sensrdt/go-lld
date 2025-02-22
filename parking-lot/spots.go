package main

type Spot struct {
	SpotNumber    int
	ParkedVehicle *Vehicle
}

// Intialize a new spot
func CreateNewSpot(spotNumber int) *Spot {
	return &Spot{
		SpotNumber: spotNumber,
	}
}

func (s *Spot) IsAvailable() bool {
	// fmt.Println(s.ParkedVehicle)
	return s.ParkedVehicle == nil
}

func (s *Spot) ParkVehicle(v *Vehicle) {
	if s.IsAvailable() {
		s.ParkedVehicle = v
	}
}

func (s *Spot) UnparkVehicle(v *Vehicle) {
	s.ParkedVehicle = nil
}
