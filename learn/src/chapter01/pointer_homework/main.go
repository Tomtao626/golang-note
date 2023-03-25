package main

import "fmt"

//写一个程序，获取一个 int 变量 num 的地址，并显示到终端
//将 num 的地址赋给指针 ptr , 并通过 ptr 去修改 num 的值

func main() {
	var num int = 1
	fmt.Println(num)
	var ptr *int = &num
	fmt.Printf("ptr指向的地址%v\n", *ptr)
	*ptr = 200
	fmt.Printf("ptr指向的地址的值%v\n", *ptr)

	// 错误示范
	/*
		符号缺失 格式错误
		var a int = 300
		var ptrs *int = a 错误
	*/
	/*
		数值类型不匹配
		var a int = 300
		var ptr *float32 = &a 错误 数值类型不匹配
	*/
}
