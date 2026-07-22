package main

import (
	"fmt"
	"sync"
)

func ppworkwe(logger func()) {
	logger()
}

func PingPongHandler() {
	pingSem := make(chan struct{})
	pongSem := make(chan struct{})
	nums := 10
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < nums; i++ {
			<-pongSem
			ppworkwe(func() { fmt.Println("pong") })
			if i != nums-1 {
				pingSem <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < nums; i++ {
			<-pingSem
			ppworkwe(func() { fmt.Println("ping") })
			pongSem <- struct{}{}
		}
	}()

	pingSem <- struct{}{}

	wg.Wait()
	close(pingSem)
	close(pongSem)

}
