# Token Bucket Rate Limiter

This folder implements a production-grade rate limiter using the **Token Bucket Algorithm**.

## File Implementations

- **`main.go`**: Implements the `TokenBucket` struct which tracks `tokens` and `capacity`. A background goroutine continually refills tokens at a set rate using a `time.Ticker`. The `Allow()` method checks if tokens are available, consuming one per request and permitting bursts up to the capacity limit.
