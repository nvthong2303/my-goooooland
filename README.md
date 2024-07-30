# My Goland

## Goroutine

- Là một trong những tính năng đặc biệt nhất trong Golang, dùng trong lập trình concurrency cực kỳ đơn giản.
- Bản chất là các function / method được thực thi một cách độc lập và đồng thời nhưng vẫn có thể kết nối với nhau.
- Trong bất kỳ chương trình go nào đều tồn tại ít nhất 1 goroutine, Main Goroutine, nếu nó kết thúc thì toàn bộ các goroutine khác cũng kết thúc.
- Bản chất nó là 1 lightweight execution thread (luồng thực thi gọn nhẹ). Vì thế sử dụng các Goroutine có chi phí tài nguyên thấp hơn nhiều so với các Thread OS truyền thống (2Kb so với 2Mb).

### Khai báo và sử dụng goroutine
### Sử dụng goroutine với anonymous function
### Vấn đề capture variable trong goroutine
### Sử dụng Gosched() để force schedule Goroutines.

### Go runtime
- Cách Go runtime quản lý các goroutines và các hành vi lập lịch của hệ điều hành:

1. Scheduling của Goroutines:
    - Go runtime sử dụng một runtime scheduler để quản lý các goroutines. Scheduler này cố gắng cung cấp công bằng giữa các goroutines, nhưng không đảm bảo rằng các goroutines sẽ xen kẽ một cách hoàn hảo.
    - Khi một goroutine bắt đầu thực thi, nó có thể chiếm CPU một khoảng thời gian trước khi scheduler quyết định chuyển sang goroutine khác. Do đó, bạn có thể thấy nhiều lần "I am Goroutine 1" hoặc "I am Goroutine 2" in ra liên tục.

2. Buffering và Context Switching:
    - Việc chuyển đổi giữa các goroutines (context switching) có thể không diễn ra ngay lập tức sau mỗi lần in ra một dòng. Điều này dẫn đến việc một goroutine có thể in ra nhiều dòng liên tiếp trước khi chuyển sang goroutine khác.
    - Buffer của console output cũng có thể gây ra hiện tượng này. Console output có thể được buffer và flush một cách không đồng bộ, dẫn đến các chuỗi thông điệp liên tục từ một goroutine trước khi chuyển sang thông điệp từ goroutine khác.

### Go Scheduler:
- Khi một ứng dụng Go start, nó sẽ được cung cấp một bộ Vi xử lý ảo (Logical virtual processor - P) cho mỗi virtual core trên máy host, Nếu có nhiều hardware thread thì trên mỗi physical core (Hyper-Threading) Go sẽ koi mỗi hardware thread là một virtual core.
    ```
    fmt.Println(runtime.NumCPU())
    ```
- exec đoạn code trên, kết quả ra được là 8.
- Tức là mọi chương trình khi chạy đều được cung cấp 8 P, mỗi P sẽ được gán cho một OS Thread. Các Thread này được quản lý bởi OS và OS chịu trách nhiệm đặt các Thread này lên core, có nghĩa là chúng ta sẽ có 8 Thread sẵn sàng để thực thi công việc.
- Mọi chương trình Go cũng đều được cung cấp một Goroutine ban đầu, Goroutine giống như một OS Thread nhưng ở tầng ứng dụng, cũng như OS Thread được chuyển đổi context trên một core, các Goroutine cũng chuyển đổi context trên M (machine)
- Có 2 Run-queue khác nhau trong GoScheduler: **Global Run Queue (GRQ)** và **Local Run Queue (LRQ)**. Mỗi P được cung cấp một LRQ để quản lý các Goroutine đã được chỉ định để thực hiện bên trong context của P. Những Goroutine này sẽ lần lượt switch-context on hoặc off trên P. GRQ dành cho các Goroutine chưa được gán cho một P nào.
![image info](./z_img_002.png)

#### Cooperating Scheduler
- OS scheduler là một Preemptive Scheduler. Về cơ bản thì điều này nghĩa là không thể dự đoán được scheduler sẽ làm gì tiếp theo. Kernel sẽ chịu trách nhiệm này và mọi thứ sẽ không thể xác định rõ ràng. Các ứng dụng chạy trên OS không có quyền kiểm soát những gì xảy ra bên trong Kernel bằng scheduler ngoại trừ khi sử dụng atomic hoặc mutex.
- GoScheduler là một phần của Go Runtime, được tích hợp vào ứng dụng golang. Có nghĩa là Go Scheduler chạy trong User Space, tầng phía trên của Kernel. Hiện tại cách triển khai Go Scheduler không phải là Preemptive scheduler mà là Cooperating scheduler (scheduler này cần các sự kiện User space được xác định rõ ràng để có thể đưa ra quyết định schedule).
- ..

## Channel
-