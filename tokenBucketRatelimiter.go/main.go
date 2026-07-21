package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	bucket chan struct{}
	count  int
	n      int
}

func (r *RateLimiter) refillBucket() {
	ticker := time.NewTicker(time.Second / time.Duration(r.n))
	for range ticker.C {
		select {
		case r.bucket <- struct{}{}:
			// added
		default:
			// done
		}
	}
}

func NewRateLimiter(n int) *RateLimiter {
	rateLimiter := &RateLimiter{
		bucket: make(chan struct{}, n),
		count:  0,
		n:      n,
	}

	for range n {
		rateLimiter.bucket <- struct{}{}
	}

	return rateLimiter
}

func (r *RateLimiter) CanTake() bool {
	select {
	case <-r.bucket:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Take() {
	<-r.bucket
	r.count++
}

func main() {

	rl := NewRateLimiter(1) // 2 operations per second

	go rl.refillBucket()

	// First two succeed immediately.
	fmt.Println(rl.CanTake()) // true
	fmt.Println(rl.CanTake()) // true

	// Bucket is empty.
	fmt.Println(rl.CanTake()) // false

	fmt.Println("Waiting for tokens...")

	rl.Take()
	fmt.Println("Allowed at:", time.Now().Format("15:04:05.000"))

	rl.Take()
	fmt.Println("Allowed at:", time.Now().Format("15:04:05.000"))

	rl.Take()
	fmt.Println("Allowed at:", time.Now().Format("15:04:05.000"))

}
