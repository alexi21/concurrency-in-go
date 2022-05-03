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

func Fanout(
	done <-chan interface{},
	input <-chan int,
	outputOne chan int,
	outputTwo chan int,
) {
	go func() {
		for {
			select {
			case outputOne = <-input:
			case outputTwo = <-input:
			case <-done:
				return
			}
		}
	}()
}

func main() {
	done := make(chan interface{})
	stream := createStream_(done)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-stream)
	}
	close(done)
	<-stream
}
