package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	bucket chan struct{}
	rate   int
}

// Create a new rate limiter.
func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		bucket: make(chan struct{}, rate),
		rate:   rate,
	}

	// Fill the bucket initially.
	for i := 0; i < rate; i++ {
		rl.bucket <- struct{}{}
	}

	// Start refilling tokens.
	go rl.refill()

	return rl
}

// Add one token every (1/rate) second.
func (r *RateLimiter) refill() {
	ticker := time.NewTicker(time.Second / time.Duration(r.rate))
	defer ticker.Stop()

	for range ticker.C {
		select {
		case r.bucket <- struct{}{}:
			// Added a token.
		default:
			// Bucket already full.
		}
	}
}

// Blocking call.
// Waits until a token is available.
func (r *RateLimiter) Take() {
	<-r.bucket
}

// Non-blocking call.
// Returns false immediately if no token is available.
func (r *RateLimiter) TryTake() bool {
	select {
	case <-r.bucket:
		return true
	default:
		return false
	}
}

func main() {
	// Allow 3 requests per second.
	rl := NewRateLimiter(4)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			rl.Take() // Wait for permission.
			// <-rl.bucket

			fmt.Printf("Request %d executed at %s\n",
				id,
				time.Now().Format("15:04:05.000"))
		}(i)
	}

	wg.Wait()
}
