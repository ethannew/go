package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	// producer
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ch <- i
		}

		defer close(ch)
	}()

	// consumer
	for item := range ch {
		fmt.Printf("received %d\n", item)
	}
}
