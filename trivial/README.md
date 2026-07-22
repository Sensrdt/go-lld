# Trivial Puzzles & Brain Teasers

This folder contains small, tricky coding puzzles often asked in Go interviews to test deep language knowledge.

## File Implementations

- **`alphabate.go`**: Prints the alphabet asynchronously using `sync.WaitGroup` and goroutines, demonstrating how to avoid common closure and deadlock pitfalls.
- **`fizzbuzz.go`**: A concurrent FizzBuzz implementation using `sync.Mutex` and `sync.Cond` to orchestrate exactly when multiple worker threads print "Fizz", "Buzz", or numbers.
- **`map.go`**: Demonstrates executing a map function in parallel over a slice using goroutines.
- **`pingpong.go`**: Implements two goroutines passing a token back and forth ("Ping" and "Pong") via unbuffered channels.
- **`pipeline.go`**: Builds a concurrent data processing pipeline (generator -> processor -> consumer) using channels.
- **`ringBuffer.go` & `ringOfChannels.go`**: Implements circular and chained communication patterns where goroutines pass signals in a ring formation.
- **`main.go`**: Runner file to execute any of the above handlers.
