package main

import (
	"fmt"
	"sync"
)

type FizzBuzz struct {
	mu   sync.Mutex
	cond sync.Cond
	cur  int
	n    int
}

func Constructor(n int) *FizzBuzz {

	fb := &FizzBuzz{
		mu:  sync.Mutex{},
		cur: 1,
		n:   n,
	}

	fb.cond = *sync.NewCond(&fb.mu)

	return fb
}

func (fb *FizzBuzz) Worker(logic func(int) bool, runner func(int)) {

	for {
		fb.mu.Lock()
		for !logic(fb.cur) && fb.cur <= fb.n {
			fb.cond.Wait()
		}

		if fb.cur > fb.n {
			fb.mu.Unlock()
			return
		}

		runner(fb.cur)
		fb.cur++
		fb.mu.Unlock()
		fb.cond.Broadcast()
	}

}

func FizzBuzzHandler() {

	fb := Constructor(15)

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		fb.Worker(func(i int) bool {
			return i%3 == 0 && i%5 == 0
		}, func(i int) {
			fmt.Printf("FizzBuzz")
		})
	}()
	go func() {
		defer wg.Done()
		fb.Worker(func(i int) bool { return i%3 == 0 && i%5 != 0 }, func(i int) { fmt.Printf("Fizz") })
	}()
	go func() {
		defer wg.Done()
		fb.Worker(func(i int) bool { return i%3 != 0 && i%5 == 0 }, func(i int) { fmt.Printf("Buzz") })
	}()
	go func() {
		defer wg.Done()
		fb.Worker(func(i int) bool { return i%3 != 0 && i%5 != 0 }, func(i int) { fmt.Printf("%d", i) })
	}()

	wg.Wait()
}
