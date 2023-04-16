package main

import "fmt"

func main() {

	// 第一种匿名函数
	func(){
		fmt.Println("=====这是匿名函数")
	}()

	// 第二种匿名函数
	f := func(){
		fmt.Println("+++++这是第二种匿名函数")
	}
	f()

	// 扩展
	ret := func(x,y int) int{ // 形参
		fmt.Println("=====这是匿名函数")

		fmt.Println(x + y)
		return x + y
	}(3,4)  // 实参

	fmt.Println(ret)



}
