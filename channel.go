package main

import (
	"fmt"
	"time"
)

func main() {
	// myChan := make(chan int)
	// go func() {
	// 	myChan <- 1
	// }()
	// fmt.Println(<-myChan)

	myChan := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChan <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(<-myChan)
	}
}