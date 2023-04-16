package main

import "fmt"

func main() {

	a := new(int)

	fmt.Println(a)

	// panic: runtime error: invalid memory address or nil pointer dereference
	var x *int // x指向内存地址，内存地址不能直接赋值

	x = new(int)
	*x = 10
	fmt.Println(x)
	fmt.Println(*x)

}
