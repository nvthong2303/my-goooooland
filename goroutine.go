package main

import (
	"fmt"
	"time"
	"runtime"
)

func name(name string) {
	for i := 0; i <= 5; i++ {
		fmt.Printf("Hello %s\n", name)
	}
	time.Sleep(time.Second)
}

func main() {
	// khai báo và sử dụng goroutine
	// go name("thong")
	// name("thuong")
	// => only print "Hello thuong"
	// go name() sẽ trả về ngay lập tức chứ k execute ngay
	// vì main kết thúc trước khi goroutine chạy xong 
	// nên các goroutine sau đó chưa đc execute
	// thêm sleep để chờ goroutine chạy xong
	// time.Sleep(time.Second)
	// => sẽ có trường hợp có in ra Hello Thong đủ hoặc ít hơn 5 lần, hoặc không lần nào
	// vì không có gì đảm bảo rằng goroutine sẽ chạy xong sau 1s trước khi main kết thúc.
	// phụ thuộc vào runtime 

	// sử dụng goroutine anonymous function
	// go func() {
	// 	for i := 0; i <= 5; i++ {
	// 		fmt.Println(i)
	// 	}
	// } ()
	// time.Sleep(time.Second)

	// vấn đề capture variable trong goroutine
	// for i := 1; i <= 100; i++ {
	// 	go func() {
	// 		fmt.Println(i) // biến i ở đây là một pointer
	// 	}()
	// }
	// time.Sleep(time.Second)
	// có thấy các giá trị i bị trùng lặp hoặc không được in ra và in ra không đúng thứ tự.
	// nếu một hàm sử dụng biến ở ngoài nó, nó sẽ được capture, có thể hiểu nó chỉ là một tham chiếu
	// vì các goroutine không được thực thi ngay thời điểm đó nên khi thực thi, giá trị i đã bị thay đổi.
	// giải pháp:
	// for i := 1; i <= 100; i++ {
	// 	go func(value int) {
	// 		fmt.Println(value) // value ở đây độc lập với i ở ngoài
	// 	}(i) // value i được copy ở đây
	// }
	// time.Sleep(time.Second)
	// => in ra các giá trị từ 1 đến 100 ngẫu nhiên và không trùng nhau.

	// sử dụng Gosched() để force schedule Goroutines.
	// go func() {
	// 	for i := 1; i <= 50; i++ {
	// 		fmt.Println("I am Goroutine 1")
	// 	}
	// }()
	// go func() {
	// 	for i := 1; i <= 50; i++ {
	// 		fmt.Println("I am Goroutine 2")
	// 	}
	// }()
	// time.Sleep(time.Second)
	// khả năng cao "I am Goroutine 1" hoặc "I am Goroutine 2" in ra liên tục mới đến cái còn lại.
	// Go runtime sử dụng một runtime scheduler để quản lý các goroutines. 
	// Scheduler này cố gắng cung cấp công bằng giữa các goroutines, 
	// nhưng không đảm bảo rằng các goroutines sẽ xen kẽ một cách hoàn hảo.
	// Khi một goroutine bắt đầu thực thi, nó có thể chiếm CPU một khoảng thời gian 
	// trước khi scheduler quyết định chuyển sang goroutine khác. 
	// Do đó, bạn có thể thấy nhiều lần "I am Goroutine 1" hoặc "I am Goroutine 2" in ra liên tục.
	// Giải pháp: sử dụng Gosched():
	// go func() {
	// 	for i := 1; i <= 50; i++ {
	// 		fmt.Println("I am Goroutine 1")
	// 		runtime.Gosched()
	// 	}
	// }()
	// go func() {
	// 	for i := 1; i <= 50; i++ {
	// 		fmt.Println("I am Goroutine 2")
	// 		runtime.Gosched()
	// 	}
	// }()
	// time.Sleep(time.Second)
	// Gosched() chỉ giúp chúng ta schedule goroutines. 
	// Việc chọn và thực thi Goroutine nào tiếp theo là do runtime quyết định. 
	// Cách này sẽ rât hiệu quả nếu bạn có rất nhiều Goroutines trong chương trình đang làm việc với vòng lặp.

	// fmt.Println(runtime.NumCPU())

}