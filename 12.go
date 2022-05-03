package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan interface{})

	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	<-done
}
