package main

import (
	"fmt"
	"sync"
)

func RingOfChannelHandler() {
	/*
		A B C A B C A B C.....
	*/
	nums := 10
	var sem [10]chan struct{}
	for i := 0; i < nums; i++ {
		sem[i] = make(chan struct{})
	}
	wg := sync.WaitGroup{}
	for i := 0; i < nums; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-sem[i]

			fmt.Println(string('A' + (i % 3)))

			if i != nums-1 {
				sem[i+1] <- struct{}{}
			}

		}(i)
	}

	sem[0] <- struct{}{}
	wg.Wait()
}
