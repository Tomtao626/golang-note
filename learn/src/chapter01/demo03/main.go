package main

import "fmt"

// 定义全局变量
var n1 = 100
var n2 = 200
var name1 = "jack"

// 上面的声明方式等价于以下写法 一次性声明
var (
	n3    = 100
	n4    = 200
	name2 = "mary"
)

func main() {
	//一次性声明多个变量1
	//var a, b, c int
	//fmt.Println(a, b, c)
	//一次性声明多个变量2
	//var a ,b, c = 100, "tom_tao", 10.11
	//fmt.Println(a, b, c)
	//一次性声明多个变量3 类型推导
	a, b, c := 100, "tom_tao", 10.11
	fmt.Println(a, b, c)

	// 输出全局变量
	fmt.Println(n1, n2, name1)
	fmt.Println(n3, n4, name2)
}
