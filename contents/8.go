package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	t := time.Since(start).Round(time.Millisecond)

	ch <- result{url, nil, t}
}

var list = []string{
	"https://amazon.com",
	"https://google.com",
	"https://nytimes.com",
	"https://wsj.com",
}

func main() {
	// After waits for the duration to elapse
	// and then sends the current time on the returned channel
	stopper := time.After(5 * time.Second)
	results := make(chan result)
	t := time.Now()

	for _, url := range list {
		go get(url, results)
	}

Loop:
	for {
		select {
		case r := <-results:
			log.Printf("%-20s %s\n", r.url, r.latency)
		case t := <-stopper:
			log.Fatalf("timeout %s", t)
			break Loop
		}
	}
	log.Println("Total time: ", time.Since(t).Round(time.Millisecond))
}
