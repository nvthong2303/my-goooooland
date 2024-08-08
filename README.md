# My Goland
Không có gì đặc biệt, nơi tôi tự học golang thôi :D

## Systax / Feature:
### Datatype:
- Basic type: numbers, strings, boolean.
    - numbers:
    - string:
    - boolean:
    - byte:
    - rune: 

- Aggregate type (kiểu dữ liệu tham chiếu): array, structs
    - arr: mảng có kích thước cố định, không thể thay đổi sau khi khai báo
    - struct: ~object(js) hoặc dict(py), các trường không được truyền giá trị có giá trị mặc định là 0.
- Reference type (kiểu tổng hợp): pointers, slices, maps, functions, channels
    - slices: kiểu dữ liệu tổng hợp, biểu diễn một dymanic array, chúng giống array nhưng kích thước có thể thay đổi trong khi chương trình chạy và chúng tham chiếu đến dữ liệu chứ không phải giá trị.
    - maps: (kiểu dữ liệu tham chiếu) tập hợp các cặp key-value không có thứ tự, các key phải unique nhưng value có thể lặp lại.
    - pointers: 
    - interfaces: 
- Interface

### Generics:
- hiểu đơn giản, generics là 1 khái niệm có trong nhiều ngôn ngữ, generics programming là việc định nghĩa các function mà không định nghĩa kiểu dữ liệu sẽ sử dụng hoặc trả về. 

## 1. Goroutine

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

## 2. Channel
- Là các kênh giao tiếp trung gian giữa các Goroutine trong Golang. Giúp các goroutine có thể gửi và nhận dữ liệu một cách an toàn thông qua cơ chế lock-free.
- Mặc định, giao tiếp trong channel là giao tiếp 2 chiều, nghĩa là channel có thể dùng cho cả gửi và nhận dữ liệu.
- Việc gửi dữ liệu vào channel sẽ giống như "Tôi đã hoàn tất công việc của mình với dữ liệu này và bàn giao chúng cho người khác".
- Cơ chế Block của channel: Việc gửi và nhận qua channel có hỗ trợ cơ chế Block, việc này giúp các Goroutines giao tiếp qua Channel một cách đồng bộ. Về nguyên tắc, Channel sẽ blocl goroutines nếu nó chưa sẵn sàng (Deadlock).
- Có thể định nghĩa các func sử dụng channel chỉ nhận hoặc gửi dữ liệu.
- Close() : dùng để đóng channel, không thể gửi data vào channel đã bị đóng.
- Có thể for để lặp qua các giá trị được gửi vào 1 channel, vòng lặp tự động kết thúc khi goroutine gửi tín hiệu close channel.
    ```
    // s là 1 channel chỉ gửi
    for n := range s {
        fmt.Println("n: ", n)
        counter += n
    }
    ```

- Channel (UnBuffered Channel), là khi 1 goroutine A gửi dữ liệu đến thì nó sẽ block A lại cho đến khi có bất kỳ goroutine nào khác đến lấy dữ liệu. 

## 3. Buffered Channel
- Buffered channel là một channel có khả năng lưu trữ dữ liệu bên trong đó. Ngược lại UnBuffered Channel, nó mang trong mình 1 sức chứa (capacity). Buffered channel sẽ không bị block goroutine nếu sức chứa vẫn còn, không cần phải có 1 goroutine khác đến lấy dữ liệu. Nó sẽ block goroutine hiện tại nếu vượt quá sức chứa.
- Buffered channel có 2 thuộc tính len (số lượng dữ liệu đang có trong bufferd channel) và cap (sức chứa tối đa).
- Đọc dữ liệu từ Buffered channel sẽ block goroutine (giống unBuffered Channel).
- Lưu trữ dữ liệu theo FIFO.



## 4. Panic and Recover, Defer:
- Đây là 2 Tính năng (Panic và Recover) được sử dụng để quản lý lỗi trong chtrinh Go.
- Panic: tương tự như execption trong js hay python (nó là 1 exeption trong go). Panic được gây ra bởi 1 lỗi runtime và gọi thẳng đến hàm Panic trong go (tích hợp sẵn). Panic có thể xảy ra theo 2 cách: Lỗi runtime của chtrnh hoặc được gọi trực tiếp.
    - Panic lỗi runtime: khi chtrnh gặp lỗi (truy cập index quá giới hạn của arr, gửi data vào channel đã đóng, ... ) Panic được tạo ra, bao gồm 2 điều:
        - thông báo lỗi
            exp: 
            ```
            panic: runtime error: index out of range [2] with length 2
            ```
        - trace của ngăn xếp nơi xảy ra panic
            exp: 
            ```
            goroutine 1 [running]:
            main.checkAndPrint(...)
                    main.go:12
            main.main()
                    /main.go:8 +0x1b
            exit status 2
            ```
    - Panic được gọi trực tiếp: dùng để bắt các ngoại lệ (input không hợp lệ khiến chtrinh không thể tiếp tục, ...)

- Defer: Khi Panic được kích hoạt, hàm đang thực thi sẽ dừng lại và các hàm Defer (trong ngăn xếp) được gọi cho đến khi tất cả chúng được trả về. Lúc đó chtrnh mới dừng và trả ra Panic. ([Tham khảo](https://tuhocweb.com/golang-nang-cao-panic-va-recover-trong-golang-156.html))

- Recover: hàm được tích hợp sẵn trong go, dùng để lấy lại quyền kiểm soát goroutine đang panic. Hàm Recover trả về giá trị được truyền cho hàm panic và không bị Side Effect, nghĩa là goroutine không bị panic, hàm recover sẽ trả về  nil . Do đó việc kiểm tra giá trị trả về của recover có nil không để biết chtrinh có đang bị panic hay không.
    +  hàm defer() là hàm duy nhất được phép thực thi sau khi panic xảy ra, vì vậy đặt recover trong hàm defer là hợp lý nhất. Nếu không sẽ không ngăn chặn được panic.
    + Recover chỉ có thể khôi phục lỗi panic xảy ra trong cùng 1 goroutine.
