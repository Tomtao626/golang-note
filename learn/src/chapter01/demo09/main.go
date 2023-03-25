package main

import "fmt"

func main() {
	// 求两个数的最大值
	x := 10
	y := 15
	var max int
	var min int
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}
	fmt.Printf("max:%d, min:%d\n", max, min)

	// 求三个数中的最大值

	var z = 20
	if z > max {
		max = z
	}
	fmt.Printf("max:%d\n", max)

	// 传统的三元运算符  n = i > j ? i : j golang不支持
	// golang本身不支持三元运算符
	var n int
	var i int = 9
	var j int = 10
	if i > j {
		n = i
	} else {
		n = j
	}
	fmt.Println("n=", n)
}
