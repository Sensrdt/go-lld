package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func isPrime(number int, primeChannel chan<- int, nonPrimeChannel chan<- int) {
	if number < 2 {
		nonPrimeChannel <- number
		return
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			nonPrimeChannel <- number
			return
		}
	}
	primeChannel <- number
}

func PrimeChecker() {
	numbers := []int{2, 4, 7, 9, 11, 15, 17, 21, 23, 50, 123, 12, 45, 32, 24}
	var wg sync.WaitGroup
	primeChannel := make(chan int, len(numbers))
	nonPrimeChannel := make(chan int, len(numbers))

	for _, val := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			isPrime(num, primeChannel, nonPrimeChannel)
		}(val)
	}

	go func() {
		wg.Wait()
		close(primeChannel)
		close(nonPrimeChannel)
	}()

	for {
		select {
		case v, ok := <-primeChannel:
			if !ok {
				primeChannel = nil
			} else {
				fmt.Println("Prime: ", v)
			}
		case v, ok := <-nonPrimeChannel:
			if !ok {
				nonPrimeChannel = nil
			} else {
				fmt.Println("Non Prime: ", v)
			}
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout reached")
			return
		}

		// Exit when both channels are closed
		if primeChannel == nil && nonPrimeChannel == nil {
			break
		}
	}
}
