package main

import (
	"fmt"
	"sync"
)

func RingBufferHandler() {
	/*
		G1
		G2
		G3
		G4
		G5
		G1
		G2
		G3
		G4
		G5
	*/

	wg := sync.WaitGroup{}
	const nums = 5
	var sem [nums]chan struct{}
	for i := 0; i < nums; i++ {
		sem[i] = make(chan struct{})
	}
	const r = 2
	for i := 0; i < nums; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < r; j++ {
				<-sem[i]
				fmt.Printf("G%d ", i)
				isLastTurn := j == r-1 && i == nums-1
				if !isLastTurn {
					sem[(i+1)%nums] <- struct{}{}
				}
			}
		}(i)
	}

	sem[0] <- struct{}{}
	wg.Wait()
}
