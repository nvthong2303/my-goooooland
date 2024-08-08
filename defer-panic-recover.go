package main

import (
	"fmt"
	"time"
)

func main() {
	// =================== PANIC ===================
	// panic runtime
	// a := []string{"a", "b"}
	// print(a, 2)
	// => 	panic: runtime error: index out of range [2] with length 2

	// panic được gọi trực tiếp
	// a := []string{"a", "b"}
	// checkAndPrint(a, 2)
	// => 	panic: Out of bound access for slice

	// =================== DEFER ===================
	// defer fmt.Println("defer 1")
	// panic("panic with defer")
	// fmt.Println("after panic")
	// => 	đến panic là dừng, không in ra "after panic"
	// => 	defer 1
	// 		panic: panic with defer

	// defer nhiều lớp
	// f1()
	// => 	defer f2
	// 		defer f1
	// 		panic: panic in f2

	// =================== RECOVER ===================
	// a := []string{"a", "b"}
	// checkAndPrint_r(a, 2)
	// fmt.Println("Exiting normally")
	// => 	Recovering from panic: Out of bound access for slice
	// 		Exiting normally

	// a := []string{"a", "b"}
    // checkAndPrintWithRecover(a, 2)
    // fmt.Println("Exiting normally")


	// recover va panic khac goroutine
	a := []string{"a", "b"}
    checkAndPrintWithRecover_c(a, 2)
    time.Sleep(time.Second)
    fmt.Println("Exiting normally")
}

func checkAndPrintWithRecover_c(a []string, index int) {
    defer handleOutOfBounds_c()
    go checkAndPrint_c(a, 2)
}
func checkAndPrint_c(a []string, index int) {
    if index > (len(a) - 1) {
        panic("Out of bound access for slice")
    }
    fmt.Println(a[index])
}
func handleOutOfBounds_c() {
    if r := recover(); r != nil {
        fmt.Println("Recovering from panic:", r)
    }
}

func print(a []string, index int) {
	fmt.Println(a[index])
}
func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}
func f1() {
	defer fmt.Println("defer f1")
	f2()
	fmt.Println("after panic in f1")
}
func f2() {
	defer fmt.Println("defer f2")
	panic("panic in f2")
	fmt.Println("after panic in f2")
}
func checkAndPrint_r(a []string, index int) {
	defer handleOutOfBounds()
	// hàm recover trên sẽ bắt đc panic và in ra thông báo
	// sau khi recover() đc gọi thì chtrinh tiếp tục chạy
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}
func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
func checkAndPrintWithRecover(a []string, index int) {
    defer handleOutOfBounds_s()
    checkAndPrint_s(a, 2)
}
func checkAndPrint_s(a []string, index int) {
    if index > (len(a) - 1) {
        panic("Out of bound access for slice")
    }
    fmt.Println(a[index])
}
func handleOutOfBounds_s() {
    if r := recover(); r != nil {
        fmt.Println("Recovering from panic:", r)
    }
}
func checkAndPrint_t(a []string, index int) {
	handleOutOfBounds_t()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}
func handleOutOfBounds_t() {
	// recover không nằm trong defer, main sẽ dừng khi gặp panic
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

