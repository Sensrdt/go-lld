package main

import (
	"fmt"
	"sync"
	"time"
)

func runner(nums []int, wg *sync.WaitGroup, logic func(int) int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			updatedValue := logic(nums[i])
			// println(updatedValue)
			result[i] = updatedValue
		}(i)
	}
	return result
}

type ArrObj struct {
	index int
	value int
}

func worker(i int, value <-chan ArrObj, result chan<- ArrObj, wg *sync.WaitGroup, logic func(int) int) {
	for v := range value {
		response := logic(v.value)
		result <- ArrObj{
			value: response,
			index: v.index,
		}
	}
}

func ParallelMap() {
	wg := sync.WaitGroup{}
	startTime := time.Now()
	result := runner([]int{1, 2, 3, 5, 6, 7}, &wg, func(i int) int {
		return i + i
	})

	fmt.Println(result)

	/*
		workers := 3
		value := make(chan ArrObj)
		result := make(chan ArrObj)
		logic := func(i int) int {
			return i * 1
		}
		for i := range workers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				worker(i, value, result, &wg, logic)
			}()
		}

		nums := 1000
		go func(nums int) {
			defer close(value)
			for i := 0; i < nums; i++ {
				value <- ArrObj{
					index: i,
					value: i,
				}
			}
		}(nums)

		go func() {
			wg.Wait()
			close(result)
		}()

		ans := make([]int, nums)
		for finalValue := range result {
			ans[finalValue.index] = finalValue.value
		}

		fmt.Println(ans)
	*/

	fmt.Println(time.Since(startTime))

}
