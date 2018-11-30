package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			out = nil
			fmt.Println("producer exit")
		}()

		for i := 0; i < n; i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Millisecond)
		}
	}()
	return out
}

func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
		}()

		// Using for-range to exit goroutine
		for x := range in {
			fmt.Printf("Process %d\n", x)
		}
	}()

	return finish
}

func main() {
	out := producer(3)
	finish := consumer(out)

	// Wait consumer exit
	<-finish
	fmt.Println("main exit")
}
