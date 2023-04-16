package main

import "fmt"

func Add(x,y int,f func(a,b int))  {
	f(x,y) // 实参

}

func main() {

	// 匿名函数作为回调函数

	Add(4,5, func(a, b int) { // 形参
		fmt.Println(a+b)

	})

}
