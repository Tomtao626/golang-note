package main

import "fmt"

func swap(n1 *int, n2 *int) {
	*n1, *n2 = *n2, *n1
}

// 请编写一个函数swap(n1 *int, n2 *int) 交换n1和n2的值
func main() {
	a, b := 10, 20
	swap(&a, &b) //传入的是地址
	fmt.Printf("a=%v,b=%v\n", a, b)
	fmt.Println("hello world")
}
