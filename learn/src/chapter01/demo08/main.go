package main

import "fmt"

/*
func main(){
	// 有两个整数a，b 交换a，b的值，最终打印结果
	a := 10
	b := 78
	fmt.Printf("交换前a和b的值分别为a=%v, b=%v\n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("交换后a和b的值分别为a=%v, b=%v\n", a, b)

	// 复合赋值的操作
	x := 11
	x += x
	fmt.Printf("x的值为：%v", x)
}
*/

func main() {
	// 有两个整数a，b 交换a，b的值，但不允许使用中间变量，最终打印结果
	a := 10
	b := 78
	fmt.Printf("交换前a和b的值分别为a=%v, b=%v\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("交换后a和b的值分别为a=%v, b=%v\n", a, b)
}
