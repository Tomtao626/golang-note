package main

import (
	"fmt"
	"strings"
)

//闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)
//累加器
func AddUpper() func(int) int {
	var i int = 1
	var str = "hello"
	return func(x int) int {
		i = i + x
		str += string(36) // 36 ascii --> $
		fmt.Println("str=", str)
		return i
	}
}

//1.编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名，如.jpg，并返回一个闭包
//2.调用闭包，可以传入一个个文件名，如果该文件名没有指定的后缀，如.jpg,则返回文件名.jpg，如果已经有.jpg后缀，则返回原文件名
//3.要求使用闭包的方式完成
//4.strings.HasSuffix 该函数可以判断某个字符串是否有指定的后缀
func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if !strings.HasSuffix(filename, suffix) {
			return filename + suffix
		}
		return filename
	}
}

//普通函数实现
func makeSuffix2(suffix string, filename string) string {
	if !strings.HasSuffix(filename, suffix) {
		return filename + suffix
	}
	return filename
}

func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) //2
	fmt.Println(f(2)) //4
	fmt.Println(f(3)) //7

	//测试makeSuffix的使用
	//返回一个闭包
	//返回的匿名函数和makeSuffix(suffix string)的suffix变量组成了一个闭包，因为返回的函数引用到了suffix这个变量
	f2 := makeSuffix(".png")
	fmt.Println("before=", f2("tom"))    //before= tom.png
	fmt.Println("after=", f2("tao.png")) //after= tao.png
	fmt.Println("after=", f2("tao.jpg")) //after= tao.jpg.png
	//普通函数调用
	fmt.Println(makeSuffix2(".jpg", "tom.pmg"))
	fmt.Println(makeSuffix2(".jpg", "tom.jpg"))

}

/*
	str= hello$
	2
	str= hello$$
	4
	str= hello$$$
	7
*/

/*
对上面的代码说明：
1.AddUpper是一个函数 返回的数据类型是func (int) int
2.闭包的说明：
	返回的是一个匿名函数，但是这个匿名函数引用到函数外的i，因此这个匿名函数和外部的i构成了一个整体，即闭包
3.也可以这样理解，闭包是类，函数是操作，n是字段。函数和它使用的i构成闭包
4.当我们反复使用闭包调用f函数时，因为i只会初始化一次，因此每调用一次就进行累计
5.闭包的关键就是：要分析出返回的函数和它使用的（引用）到哪些变量，因为函数和他引用到的变量共同构成闭包。
*/
