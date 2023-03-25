package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("n1 is :", n1)
}

func getSumandSub(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("n1 is :", n1)

	// go函数支持返回多个值
	// 如果返回多个值，在接收时，希望忽略返回某个值，则使用 _符号来代替需要忽略的返回值， 标示占位忽略
	_, res := getSumandSub(20, 10)
	fmt.Println("res is :", res)

	// 如果返回值只有一个，（返回值类型列表）可以不写()
	sum, sub := getSumandSub(7, 8)
	fmt.Printf("sum is %v\nsub is %v\n", sum, sub)
}
