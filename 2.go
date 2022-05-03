package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	
	go func() {
		c <- "Hello World"
		time.Sleep(2 * time.Second)
		c <- "I just had a quick nap"
	}()


	fmt.Println(<-c)
	fmt.Println(<-c)
}
