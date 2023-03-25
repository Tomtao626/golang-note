package main

import "fmt"

func main() {
	//第一种:指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println(i)
	//根据值自行判定变量类型(类型推导)
	var num = 10.11
	fmt.Println(num)
	//第三种:省略 var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 等价于 var name string  name = "tom_tao"
	name := "tom_tao"
	fmt.Println(name)
}
