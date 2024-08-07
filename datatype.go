package main

import (
	"fmt"
)

func main() {
	// int
	var a int = 10
	fmt.Println("int : ", a)

	//  float
	var b float64 = 3.14
	fmt.Println("float : ", b)

	// complex
	var c complex64 = 3.2 + 12i // shorthand
	var c1 complex64 = complex(3.2, 12) 
	fmt.Println("complex : ", c, c1)

	// string
	var d string = "Hello"
	fmt.Println("string : ", d)
	// string \n
	var d1 string = "\nHello\nWorld"
	fmt.Println("string has \\n: ", d1)
	// but `Hello \n World` not working :D
	// string \t
	var d2 string = "Hello\tWorld"
	fmt.Println("string has \\t: ", d2)
	// string \"
	var d3 string = "Hello\"World"
	fmt.Println("string has \\\": ", d3)
	// string \'
	var d4 string = "Hello'World"
	fmt.Println("string has \\': ", d4)
	// string backslash
	var d5 string = "Hello\\World"
	fmt.Println("string has \\\\ : ", d5)

	// boolean
	var e bool = true
	fmt.Println("boolean : ", e)

	// array
	var f [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array : ", f)
	fmt.Println("array[0] : ", f[0])

	// for loop
	fmt.Println("for loop : ")
	for i := 0; i < len(f); i++ {
		fmt.Println("array[", i, "] : ", f[i])
	}

	// slice
	var g []int = []int{}
	fmt.Println("slice : ", g)
	// make
	var g1 []int = make([]int, 5)
	fmt.Println("slice with make : ", g1)
	// make với capacity, có nghĩa slice sau ban đầu có 5 phần tử, 
	// nhưng có thể mở rộng thành 10 phần tử
	var g2 []int = make([]int, 5, 10)
	fmt.Println("slice with make and capacity : ", g2)


	// struct
	type Person struct {
		name string
		age int
		address string
	}
	var h Person = Person{
		name: "thong",
		age: 25,
		address: "Hanoi",
	}
	var h1 Person = Person{"thong", 25, "Hanoi"}
	fmt.Println("empty struct : ", Person{})
	fmt.Println("struct : ", h)
	fmt.Println("struct get value : ", h1.name)

	// map
	var i map[string]int = map[string]int{}
	fmt.Println("map : ", i)
	i["index"] = 1
	fmt.Println("map : ", i)
	fmt.Println("map : ", i["name"])
	i["banana"] = 99
	i["apple"] = 100
	i["orange"] = 101
	fmt.Println("map before delete : ", i)
	// delete
	delete(i, "banana")
	fmt.Println("map after delete : ", i)

	// pointer
	var j int = 10
	var j1 *int = &j
	fmt.Println("pointer : ", j1)
	fmt.Println("pointer get value : ", *j1)

	// interface
	type Animal interface {
		Speak() string
		Eat() string
		Write([]byte) (int, error)
	}
}