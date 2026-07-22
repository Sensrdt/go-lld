# Go Concurrency Patterns

This folder explores various concurrency constructs, patterns, and exercises in Go.

## File Implementations

- **`crawler.go`**: Implements a concurrent web crawler. Uses goroutines to fetch URLs and tracks responses (`UrlResponseRegister`) to prevent duplicate fetching.
- **`prime.go`**: A concurrent prime number checker. Uses channels (`primeChannel`, `nonPrimeChannel`) to route and filter numbers processed by worker goroutines.
- **`url.go`**: Checks the status of multiple URLs concurrently. Uses `successChannel` and `failureChannel` to aggregate `UrlResponse` structures.
- **`main.go`**: Entry point that can execute any of the concurrency handlers defined in the other files.
