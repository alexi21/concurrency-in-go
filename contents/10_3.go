package main

import "fmt"

func main() {
	c := make(chan int)

	close(c)

	i, ok := <-c
	fmt.Println(i, ok)
}
