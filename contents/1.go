package main

import (
	"fmt"
	"time"
)

func main() {

	var s string

	go func() {
		s = "Hello World"
		time.Sleep(2 * time.Second)
		s = "I just had a quick nap"
	}()


	fmt.Println("s = ", s)
	fmt.Println("s = ", s)
}
