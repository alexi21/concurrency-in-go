package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	i, ok := <-ch
	fmt.Println(i, ok)
}
