package main

import "fmt"

// 基本数据类型和数组默认都是值传递的，即进行值拷贝  在函数内修改，不会影响原来的值

func test(n1 int) {
	n1 = n1 + 10
	fmt.Println("n1=", n1)
}

func main() {
	num := 20
	test(num)
	fmt.Println("num=", num)
}
