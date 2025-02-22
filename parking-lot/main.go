package main

import "fmt"

func main() {
	fmt.Println("--------- Parking Lot LLD --------")
	pl := GetParkingLotInstance()
	pl.AddLevel(CreateNewLevel(1, 4))

	pl.DisplayParkingLotStatus()

	pl.ParkVehicle(&Vehicle{RegNumber: "KA-01-HH-1234"})
	pl.ParkVehicle(&Vehicle{RegNumber: "KA-01-HH-9999"})

	pl.DisplayParkingLotStatus()

	pl.UnparkVehicle(&Vehicle{RegNumber: "KA-01-HH-1234"})

	pl.DisplayParkingLotStatus()
}
