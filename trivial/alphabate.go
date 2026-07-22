package main

import (
	"fmt"
	"sync"
)

func Alphabet() {
	var channel [26]chan struct{}

	for i := 0; i < 26; i++ {
		channel[i] = make(chan struct{})
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 26; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-channel[i]
			fmt.Printf(string('A' + i))

			if i != 25 {
				channel[i+1] <- struct{}{}
			}
		}(i)
	}

	// start
	channel[0] <- struct{}{}
	wg.Wait()
}
