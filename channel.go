package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
	"runtime"
)

func _test() {
	time.Sleep(time.Second)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(r.Intn(100))
	wc := new(sync.WaitGroup)
	wc.Add(2)
	runtime.Gosched()
}

func receiveAndSend(c chan int) {
	fmt.Println("Receive: %d", <-c)
	fmt.Println("Send: 2 ...")
	c <- 2
}

func receiveOnly(c <-chan int) {
	fmt.Println("Receive: %d", <-c)
	// c <- 2 // error
}
func sendOnly(c chan<- int) {
	fmt.Println("Send: 2 ...")
	c <- 2
	// fmt.Printf("Received: %d\n", <-c) // error
}

func main() {
	// khai báo channel
	// gửi nhận dữ liệu qua channel
	// myChan := make(chan int)
	// go func() {
	// 	myChan <- 1
	// }()
	// fmt.Println(<-myChan)

	// có 2 goroutine, main groutine và goroutine anonymous
	// main groutine sẽ chờ goroutine anonymous chạy xong mới kết thúc
	// ở mỗi vòng lặp, main groutine với tốc độ cao hơn, sẽ phải đợi cho tới khi 
	// myChain có dữ liệu gửi vào, kết quả là in ra các số từ 1 đến 10 sau mỗi s
	// myChan := make(chan int)
	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		myChan <- i
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(<-myChan)
	// }
	// ==> cơ chế streaming data

	// Sử dụng Channel chỉ gửi và nhận data
	// myChan := make(chan int)
	// go receiveAndSend(myChan)
	// myChan <- 1
	// fmt.Println("Value from receiveAndSend: %d", <-myChan)

	// Sử dụng Channel chỉ nhận data
	// Sử dụng Channel chỉ gửi data
	
	// Đóng channel
	// myChan := make(chan int)
	// close(myChan)

	// Kiểm tra channel đã đóng chưa
	// value, isAlive := <-chanName

	// Sử dụng Select và Waitegroup
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	ch1 <- 1
	// } ()
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	ch2 <- 2
	// } ()
	// fmt.Println("Waiting for data ...")
	// fmt.Println("Data from ch1: %d", <-ch1)
	// fmt.Println("Data from ch2: %d", <-ch2)
	// // luôn đợi ch1 rồi đến ch2 măc duf cả 2 đều chạy sleep random

	// Sử dụng select để chọn Channel đã sẵn sàng:
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	ch1 <- 1
	// } ()
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	ch2 <- 2
	// } ()
	// selct {
	// 	case v1 := <-ch1:
	// 		fmt.Println("Data from ch1 come first : %d", v1)
	// 		fmt.Println("Them, data from ch2: %d", <-ch2)
	// 	case v2 := <-ch2:
	// 		fmt.Println("Data from ch2 come first : %d", v2)
	// 		fmt.Println("Them, data from ch1: %d", <-ch1)
	// }

	// sử dụng Select với for + select
	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// } ()
	// fibonacci(c, quit)
	// phương thức sử dụng 1 channel riêng để quit, thoát hàm hoặc vòng lặp
	// trong practice xuất hiện rất phổ biến, có thể xem là 1 best practice

	// Sử dụng waitgroup để biết các goroutine đã hoàn tất
	// đây sẽ là một trường hợp rất phổ biến vì chúng ra sẽ 
	// không biết khi nào các goroutine hay channel sẽ honaf tất
	// (sẽ không biết sẽ cần phải sleep trong bao lâu)
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	fmt.Println("Goroutine 1 done")
	// } ()
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	fmt.Println("Goroutine 2 done")
	// } ()
	// time.Sleep(time.Second * 6)
	// VD trên cho thấy chúng ta không biết khi nào các goroutine sẽ hoàn tất
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// wc := new(sync.WaitGroup)
	// wc.Add(2)
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	fmt.Println("Goroutine 1 done")
	// 	wc.Done()
	// } ()
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	// 	fmt.Println("Goroutine 2 done")
	// 	wc.Done()
	// } ()
	// wc.Wait()
	// fmt.Println("All goroutine done")
	// sử dụng waitgroup để biết khi nào các goroutine đã hoàn tất

	// Sử dụng channel để lắng nghe dữ liệu từ nhiều nơi
	// myChan := make(chan int)
	// go sender(myChan, "s1")
	// go sender(myChan, "s2")
	// go sender(myChan, "s3")
	// start := 0
	// for {
	// 	start += <-myChan
	// 	fmt.Println("Total: %d", start)
	// 	if start >= 300 {
	// 		break
	// 	}
	// }
	// fmt.Println("get :", <-myChan)

	unbufferedChan := make(chan int)
	unbufferedChan <- 1
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x, y = y, x + y
			case <-quit:
				fmt.Println("Quit")
				return
		}
	}
}

func sender(c chan<- int, name string) {
	for i := 1; i <= 100; i ++ {
		c <- 1
		fmt.Println("Sender %s sent: %d", name, i)
		runtime.Gosched()
	}
}