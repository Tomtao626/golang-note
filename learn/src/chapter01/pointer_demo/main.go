package main

import "fmt"

// 演示golang中的指针类型
func main() {
	//基本的数据类型
	var i int = 10
	// 输出i在内存中的地址
	fmt.Println("i的地址：", &i)
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr指向的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)
	fmt.Println(ptr)
}
