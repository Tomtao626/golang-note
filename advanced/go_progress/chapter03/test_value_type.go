package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// int, int8,int16,int32,int64

	var value1 int = 3 // 根据电脑位数

	var value2 int8 = 4

	var value3 int16 = 5

	var value4 int32 = -6

	var value5 int64 = 7


	fmt.Println(unsafe.Sizeof(value1))
	fmt.Println(unsafe.Sizeof(value2))
	fmt.Println(unsafe.Sizeof(value3))
	fmt.Println(unsafe.Sizeof(value4))
	fmt.Println(unsafe.Sizeof(value5))

	// int和uint的区别

	var value6 uint = 1
	fmt.Println(value6)

	// float32:总长是8，总长度的第九位四舍五入
	// float64:总长是16，总长度的第十七位四舍五入
	var fl1 float32 = 12233.111231531231231233123123123
	var fl2 float64 = 12233.111231231231232633123123123

	fl3 := 12233.111231531231231233123123123
	fmt.Println(fl1)
	fmt.Println(fl2)
	fmt.Println(fl3)

	// 复数
	var value7 complex64
	var value8 complex128

	fmt.Println(unsafe.Sizeof(value7))
	fmt.Println(unsafe.Sizeof(value8))


	// 整数变量和浮点型变量之间不能运算

	//a := 22
	//b := 22.22

	//c := a + b
	//fmt.Println(b + 22)



}
