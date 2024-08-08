package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	// trường hợp bình thường, chạy đúng
	// go hello()
	// time.Sleep(1 * time.Second)
	// fmt.Println("main function")

	// vấn đề waiting
	go hello()
	fmt.Println("main function")
	// main goroutine kết thúc, không đợi các goroutine khác hoàn thành.
}