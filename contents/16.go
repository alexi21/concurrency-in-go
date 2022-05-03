package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3
	}()

	time.Sleep(1 * time.Second)

	select {
	case a := <-ch1:
		fmt.Println(a)
	case a := <-ch2:
		fmt.Println(a)
	case a := <-ch3:
		fmt.Println(a)
	}
}
