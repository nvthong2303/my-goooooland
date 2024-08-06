package main

import (
	"fmt"
	"time"
)

func publisher() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("Publisher: is sending task: %d\n", i)
			c <- i
		}
		close(c)
	} ()
	return c
}

func consumer(c <-chan int, name string) {
	counter := 0
	for value := range c {
		fmt.Printf("[%s] Consumer %s: is doing task: %d\n", name, name, value)
		counter++
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Printf("[%s] Consumer %s: has finished: %d tasks\n", name, name, counter)
}

func main() {
	myChan := publisher()
	maxConsumer := 5

	for i := 1; i <= maxConsumer; i++ {
		go consumer(myChan, fmt.Sprintf("%d", i))
	}
	time.Sleep(time.Second * 10)
}