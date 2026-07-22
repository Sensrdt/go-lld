# Vending Machine Low-Level Design (LLD)

This folder contains a Low-Level Design for a standard Vending Machine using the **State Design Pattern**.

## File Implementations

- **`machine.go` & `machineOperation.go`**: `machine.go` contains the `Machine` context. `machineOperation.go` defines the `Operation` interface to which all states must conform.
- **`state.go`**: Provides concrete implementations for various machine states (Idle, HasMoney, Dispensing).
- **`inventory.go` & `products.go`**: Manage available stock, tracking quantities, and mapping to `Product` structures.
- **`coin.go` & `note.go`**: Defines the currency system, mapping constants to integer values for payment calculations.
- **`main.go`**: Simulates a user interacting with the vending machine, inserting notes, selecting a product, and dispensing change.
