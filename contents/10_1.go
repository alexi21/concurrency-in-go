package main

import "fmt"

func main() {
	c := make(chan int)

	close(c)

	fmt.Println(<-c)
}
