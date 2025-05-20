package main

import (
	"fmt"
	"time"
)

/*
	or pattern:
	Combine multiple cancellation channels into a single channel that closes when any of the input channels are closed.

	Think of it like: "As soon as any one channel is closed, we stop."
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		fmt.Println("Only one channel passed to 'or'")
		return channels[0]
	}

	done := make(chan interface{})

	go func() {
		fmt.Printf("Listening on %d channels for cancellation...\n", len(channels))
		defer close(done)

		select {
		case <-channels[0]:
			fmt.Println("channels[0] triggered cancellation")
		case <-channels[1]:
			fmt.Println("channels[1] triggered cancellation")
		case <-or(channels[2:]...): // recursion here
			fmt.Println("A deeper channel triggered cancellation (from recursion)")
		}
	}()

	return done
}

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	fmt.Println("Setting up 'or' channel to listen for any cancellation...")

	result := or(ch1, ch2, ch3)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Simulating cancellation by closing ch2...")
		close(ch2)
	}()

	select {
	case <-result:
		fmt.Println("→ 'or' channel closed: One of the input channels was closed.")
	}
}
