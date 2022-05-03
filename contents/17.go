package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
)

func main() {
	go func() {
		for {
			select {
			case ch1 <- 1:
			case ch2 <- 2:
			case ch3 <- 3:
			}
		}
	}()
	go func() {
		for {
			select {
			case a := <-ch1:
				fmt.Println(a)
			case a := <-ch2:
				fmt.Println(a)
			case a := <-ch3:
				fmt.Println(a)
			}
		}
	}()
	time.Sleep(1 * time.Millisecond)
}
