# Parking Lot Low-Level Design (LLD)

This folder contains the Low-Level Design for a scalable Parking Lot system.

## File Implementations

- **`parking_lot.go`**: The central manager (`ParkingLot` struct) coordinating the entry and exit of vehicles across multiple levels. Uses the Singleton pattern (`GetParkingLotInstance`).
- **`levels.go`**: Manages individual levels and their array of spots. Handles the logic to check if a specific level has capacity.
- **`spots.go`**: Represents a single parking space (`Spot`), tracking its availability and the vehicle currently occupying it.
- **`vehicle.go`**: Defines the `Vehicle` entity and its properties (like registration number).
- **`main.go`**: Application entry point that simulates parking operations and prints the current status of the lot.
