package main

import (
	"fmt"
	//为了使用utils.go文件中的变量或函数，我们需要先引入该model的地址
	//"golang-note/learn/chapter01/demo04/model"
)

// + 的使用
func main() {
	var a = 1
	var b = 2
	var c = a + b // 做加法运算
	fmt.Println(c)

	var j = "hello"
	var k = "world"
	var l = j + k
	fmt.Println(l)
	//	使用utils.go中的HeroName 包名.标志符
	//fmt.Println(model.HeroName)

}
