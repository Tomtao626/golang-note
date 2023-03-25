package main

import (
	"fmt"
)

var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	// 案例演示 求两个数的和 使用匿名函数的方式完成
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	//将匿名函数func (n1 int, n2 int) int 赋值给变量a
	//将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
	//则a的数据类型就是函数类型，此时，我们就可以通过a完成函数的调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 29)
	fmt.Println("res2=", res2)
	res3 := a(590, 280)
	fmt.Println("res3", res3)
	//全局匿名函数 如果将一个函数赋给一个全局变量,那么这个匿名函数就是全局匿名函数，可以在程序内有效
	res4 := Fun1(10, 20)
	fmt.Println("res4=", res4)
}
