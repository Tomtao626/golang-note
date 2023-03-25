package main

import "fmt"

func main() {
	var i int = 10
	//var a int
	//++i golang没有这种写法
	//--i golang没有这种写法
	//a = i++ golang的自增自减操作只能当作独立语言使用，不能进行赋值运算
	//a = i--
	/*
		if i--<0 {
			fmt.Println("is_ok")
		}
	*/
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
