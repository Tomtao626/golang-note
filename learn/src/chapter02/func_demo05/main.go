package main

import "fmt"

func test(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("*n1=", n1)
}

func main() {
	num := 20
	test(&num)
	fmt.Println("num=", num)
}

// 如果希望函数内的变量能够修改函数外的变量(指的默认是以值传递方式的数据类型)，可以传入变量的地址&，函数内以指针的方式操作变量，从效果上看类似引用
