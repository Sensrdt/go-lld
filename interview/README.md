# General Interview Prep

This folder contains miscellaneous scripts, exercises, and logic problems targeted specifically at cracking Go coding interviews.

## File Implementations

- **`taskProcessing.go`**: Defines a `Task` struct (unit of work) and demonstrates how to process tasks efficiently, simulating real-world worker pool interview questions.
- **`main.go`**: Contains a `SharedCounter` implementation protected by `sync.Mutex` to demonstrate safe concurrent increments, a common interview topic for handling race conditions.
