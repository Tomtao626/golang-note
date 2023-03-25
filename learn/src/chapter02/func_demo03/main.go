package main

import "fmt"

// 函数中的变量是局部的

func test() {
	//n1 是test函数的局部变量，zhin只能在test函数中使用
	var n1 int = 10
	fmt.Println(n1)
}

func main() {
	//fmt.Println("n1=", n1) # 不能使用n1 因为n1是test函数的局部变量
}
