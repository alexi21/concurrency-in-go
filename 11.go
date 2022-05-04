package main

import (
	"fmt"
	"math/rand"
)

func createStream_(done <-chan interface{}) <-chan int {
	stream := make(chan int)
	go func() {
		defer fmt.Println("createStream goroutine finished")
		defer close(stream)
		for {
			select {
			case stream <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return stream
}

func main() {
	done := make(chan interface{})
	stream := createStream_(done)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-stream)
	}
	close(done)
	//<-stream
}
