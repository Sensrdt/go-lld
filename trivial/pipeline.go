package main

import "fmt"

func ValueToChannel(arr []int) <-chan int {
	result := make(chan int)
	go func() {
		for i := 0; i < len(arr); i++ {
			result <- arr[i]
		}
		close(result)
	}()
	return result
}

func Multiplier(values <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for v := range values {
			result <- v * v
		}
		close(result)
	}()

	return result
}

func PipelineHandler() {
	arr := []int{1, 2, 4, 6, 7, 8, 32, 6, 23, 7, 1, 62, 216, 4, 123, 63}
	p1 := ValueToChannel(arr)
	p2 := Multiplier(p1)

	for v := range p2 {
		fmt.Println(v)
	}
}
