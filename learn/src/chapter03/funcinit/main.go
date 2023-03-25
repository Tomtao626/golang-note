package main

import (
	"fmt"
	//引入utils包
	"chapter03/funcinit/utils"
)

var age = test()

//为了看到全局变量是先被初始化的，先写函数
func test() int {
	fmt.Println("test()...") // 第一个执行
	return 80
}

//init函数 通常可以在init函数中完成初始化的工作
func init() {
	fmt.Println("init().....") //第二个执行
}

func main() {
	fmt.Println("main().....\nage=", age) //都三个执行
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}

//init函数的注意事项和细节
//如果一个文件同时包含变量定义，执行顺序是 变量定义>init函数>main函数
/*
test()...
init().....
main().....
age= 80
*/

// 面试题：如果main.go和utils.go中都含有变量定义,init函数时，执行的流程又是怎样的
//如果是main.go调用utils.go，会先执行utils.go中相关的操作 变量定义>init函数 ，然后再执行main.go文件中的变量定义>init函数>main函数
