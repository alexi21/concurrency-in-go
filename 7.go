package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func () {
		defer close(c)
		for i := 1; i < 10; i++ {
			c <- i
			fmt.Printf("%d write to channel\n", i)
		}
	}()

	time.Sleep(3*time.Second)
	for s := range c {
		fmt.Println(s)
		time.Sleep(1*time.Second)
	}
}
