package main

import (
	"fmt"
	"time"
)

const (
	numberOfURLs    = 10000
	numberOfWorkers = 5
)

func main() {
	// fmt.Println("Buffered Channel")
	// bufferedChan := make(chan int, 1)
	// bufferedChan <- 1
	// println(<-bufferedChan)
	queue := startQueue()
	for i := 0; i < numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i))
	}

	time.Sleep(5 * time.Second)
}

func crawlURL(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Crawl URL %d from %s\n", v, name)
		time.Sleep(time.Second)
	}
	fmt.Println("Done :", name)
}
func startQueue() <-chan int {
	queue := make(chan int)
	go func() {
		for i := 0; i < numberOfURLs; i++ {
			queue <- i
			fmt.Printf("Add URL %d\n", i)
		}
		close(queue)
	}()
	return queue
}