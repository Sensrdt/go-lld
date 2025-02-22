package main

type Levels struct {
	LevelNumber int
	Spots       []*Spot
}

func CreateNewLevel(levelNumber int, numberOfSpots int) *Levels {

	level := &Levels{LevelNumber: levelNumber}

	for i := 1; i < numberOfSpots; i++ {
		level.Spots = append(level.Spots, CreateNewSpot(i))
	}

	return level
}

func (level *Levels) ParkVehicle(v *Vehicle) bool {
	for _, spot := range level.Spots {
		if spot.IsAvailable() {
			spot.ParkVehicle(v)
			return true
		}
	}
	return false
}

func (level *Levels) UnparkVehicle(v *Vehicle) bool {
	for _, spot := range level.Spots {
		if !spot.IsAvailable() && spot.ParkedVehicle.RegNumber == v.RegNumber {
			spot.UnparkVehicle(v)
			return true
		}
	}
	return false
}

func (level *Levels) DisplayLevelStatus() {
	if len(level.Spots) == 0 {
		println("Spot is empty")
		return
	}
	for _, spot := range level.Spots {
		if spot.IsAvailable() {
			println("Spot number: ", spot.SpotNumber, " is available")
		} else {
			println("Spot number: ", spot.SpotNumber, " is not available")
		}
	}
}
