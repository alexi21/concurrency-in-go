package main

func main() {
	c := make(chan int, 1)

	<-c
}
