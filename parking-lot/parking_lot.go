package main

import "fmt"

type ParkingLot struct {
	Levels []*Levels
}

var parkingLotInstance *ParkingLot

// initializing
func GetParkingLotInstance() *ParkingLot {
	if parkingLotInstance == nil {
		parkingLotInstance = &ParkingLot{
			Levels: []*Levels{},
		}
	}

	return parkingLotInstance
}

func (pl *ParkingLot) AddLevel(level *Levels) {
	pl.Levels = append(pl.Levels, level)

}

func (pl *ParkingLot) ParkVehicle(v *Vehicle) bool {
	for _, level := range pl.Levels {
		if level.ParkVehicle(v) {
			fmt.Printf("Vehicle with registration number %s parked successfully\n", v.RegNumber)
			return true
		}
	}
	fmt.Println("Parking lot is full")
	return false
}

func (pl *ParkingLot) UnparkVehicle(v *Vehicle) bool {
	for _, level := range pl.Levels {
		if level.UnparkVehicle(v) {
			fmt.Printf("Vehicle with registration number %s unparked successfully\n", v.RegNumber)
			return true
		}
	}
	fmt.Printf("Vehicle with registration number %s not found\n", v.RegNumber)
	return false
}

func (pl *ParkingLot) DisplayParkingLotStatus() {
	if len(pl.Levels) == 0 {
		println("Level is empty")
		return
	}
	for _, level := range pl.Levels {
		level.DisplayLevelStatus()
	}
}
