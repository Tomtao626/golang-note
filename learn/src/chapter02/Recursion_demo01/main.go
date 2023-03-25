package main

import "fmt"

// 函数递归调用-一个函数在函数体内调用了本身，即称为递归调用

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n is, n: ", n)
}

func main() {

	test(4)
}
