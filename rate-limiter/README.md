# Simple Rate Limiter

This folder contains a basic interval-based rate limiter implementation in Go.

## File Implementations

- **`main.go`**: Implements a `rateLimiter` function that reads from a request channel. If the request count hits the `limit`, it forces a wait on a `time.Timer(interval)` before resetting the count. Demonstrates simple concurrency throttling.
