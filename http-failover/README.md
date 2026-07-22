# HTTP Failover (Hedged Requests)

This folder implements a highly resilient HTTP client pattern known as **Hedged Requests** or **Dynamic Failover**.

## File Implementations

- **`main.go`**: Contains the entire failover logic (`GetWithFailover`). It demonstrates launching an HTTP request and setting a 500ms `time.Timer`. If the request fails early or the 500ms elapses, it dynamically spawns the next backup request while keeping the original one running. It uses `context.WithCancel` to immediately terminate all pending requests the moment any URL returns a successful `200 OK` response.
