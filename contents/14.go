package main

import (
	"fmt"
	"time"
)

func main() {
	// Run a goroutine until main exits
	go func() {
		for i := 0; ; i++ {
			fmt.Println("Hello World: ", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Sleep for one second and then exit
	time.Sleep(1 * time.Second)
}
