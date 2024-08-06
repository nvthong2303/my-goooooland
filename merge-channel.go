package main

import (
	"fmt"
	"sync"
)

// hàm trả về 1 channel chỉ nhận dữ liệu
func streamNumbers(numbers ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, n := range numbers {
			fmt.Println("send : ", n)
			c <- n
		}
		close(c)
	} ()
	return c
}
// hàm nhận vào 1 slice các channel (chỉ nhận) và trả về 1 channel (chỉ nhận)
func sumAllStreams(streams ...<-chan int) <-chan int {
	sumChan := make(chan int)
	counter := 0
	wc := new(sync.WaitGroup)
	wc.Add(len(streams))
	for i := 0; i < len(streams); i++ {
		// mỗi channel, gọi 1 goroutine nhận vào channel (chỉ nhận) đó
		go func(s <-chan int) {
			// chỉ in ra số đầu tiên được gửi vào channel
			// fmt.Println("-------> ", <-s)
			// vòng lặp để lấy tất cả các số được gửi vào channel
			for n := range s {
				fmt.Println("receive: ", n)
				counter += n
			}
			wc.Done()
		} (streams[i])
	}

	go func() {
		fmt.Println("Run goroutine wait all")
		wc.Wait()
		sumChan <- counter
	} ()
	fmt.Println("return func sumAllStreams")
	return sumChan
}

func main() {
	s := sumAllStreams(
		streamNumbers(1, 2, 3, 4, 5),
		streamNumbers(6, 6, 7, 7, 8, 9, 10),
	)
	fmt.Println(<-s)
}



