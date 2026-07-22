# Thread-Safe Queue

This folder implements a concurrent, thread-safe Queue data structure in Go.

## File Implementations

- **`main.go`**: Contains the generic `Queue[T]` struct and its methods (`Push`, `Pop`). It uses `sync.Mutex` to lock the internal slice during mutations, preventing data races and panics when multiple goroutines attempt to modify the queue concurrently.
