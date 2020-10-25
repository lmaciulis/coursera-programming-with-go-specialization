package main

import (
	"fmt"
	"time"
)

const (
	iterationsCount = 10
)

func main() {
	num := 0

	fmt.Printf("Program started, initial value: %d\n", num)
	fmt.Print("Outputting altered integer values:\n")

	/*
		Race condition is when multiple processes trying to access and manipulate same data at a same time.
		This program is an example of race condition.
		While program runs, if every goroutine is executed in order as written in code,
		integer value always increases by 1 and then decreases by 1. So, printed runtime output should be 10101010101010101010.
		But, while running goroutines, due to race conditions, each integer addition/subtraction is unpredictable
		and every execution produces different output.
	*/
	for i := 0; i < iterationsCount; i++ {
		go addOne(&num)
		go subOne(&num)
	}

	time.Sleep(time.Millisecond * 500)
	fmt.Print("\nWait 500 ms before exiting, to ensure goroutines will finish their jobs")
}

func addOne(num *int) {
	*num++
	fmt.Print(*num)
}

func subOne(num *int) {
	*num--
	fmt.Print(*num)
}
