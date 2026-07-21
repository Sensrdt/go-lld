package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Semaphore pattern")

	maxConcurrentTask := 3
	sem := make(chan struct{}, maxConcurrentTask)

	var wg sync.WaitGroup

	wg.Add(1)

	// go func(sem chan struct{}) {
	for i := range 20 {
		go func() {
			sem <- struct{}{}

			fmt.Printf("Running task for %d \n", i)

			<-sem
		}()
	}
	// }(sem)

}
