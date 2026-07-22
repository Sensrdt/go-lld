# Interfaces in Go

This folder demonstrates the power of interfaces for building decoupled, mockable, and extensible code.

## File Implementations

- **`notificationService.go`**: Defines a `Notification` interface with a `Send` method. Implements `Email` and `SMS` structs that satisfy the interface, showing how to decouple sender logic from the transport mechanism.
- **`inMemCache.go`**: Defines a `CacheLogic` interface for key-value storage. Implements a `Cache` struct backed by a `map[string]string` to show how data stores can be abstracted.
- **`main.go`**: Entry point demonstrating the usage of both the notification and caching interfaces.
