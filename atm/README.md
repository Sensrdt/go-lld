# ATM Low-Level Design (LLD)

This folder contains a Low-Level Design (LLD) for an Automated Teller Machine (ATM) system in Go.

## File Implementations

- **`atm.go`**: Contains the core `ATM` struct which manages the state machine (using State Design Pattern) and transitions between states like `Ready`, `CardInserted`, `SelectingTransaction`, etc.
- **`bank.go`**: Implements the mock bank backend to interface with accounts, simulating network calls to verify PINs and check balances.
- **`account.go`**: Defines the `Account` entity, tracking user balances and handling validation logic.
- **`main.go`**: The entry point that wires up the ATM, the Bank, and a test Account, running through a simulated transaction flow.
