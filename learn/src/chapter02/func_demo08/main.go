package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//go支持自定义数据类型
//基础语法：
//type 自定义数据类型名 数据类型  相当于就是一个别名
// 这时myFunType就是 func(int, int) int类型
type myFunType func(int, int) int

// 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFunc2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//go支持可变参数 可变参数只能放在形参列表的最后一个 以下写法错误
//func sum(args... int, n1 int){}   error: func sum(n1 int, args... int) (sum int) {
//编写一个sum函数，可以求出1到多个int的和
//func sum(n1 int, vars... int) (sum int) {
func sum(n1 int, args ...int) (sum int) {
	sum = n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0] 表示取出args切片的一一个元素值，其他以此类推
	}
	return
}

func main() {
	res2 := myFunc(getSum, 50, 50)
	fmt.Println("res2=", res2)

	res3 := myFunc2(getSum, 50, 60)
	fmt.Printf("res3=%v", res3)

	n1, n2 := getSumAndSub(3, 4)
	fmt.Printf("\nn1=%v,n2=%v\n", n1, n2)

	//调用求和sum函数
	res4 := sum(1, 2, 23, 3, 4)
	fmt.Printf("res4=%v", res4)
}

// 函数既然是一种数据类型，因此在go中，函数可以作为形参， 并且调用
// go支持自定义数据类型， 基本语法： type 自定义数据类型 数据类型  相当于一个别名
// type mySum func(int, int) mySum等价于一个函数类型 func(int, int)

/*
type myInt int   // myInt 等价于 int  myInt和int虽然都是int类型， 但go认为myInt和int是两个数据类型

var k1 myInt
var k2 int
k1 = 40
k2 = int(k1) // 虽然都是int类型 但是需要显式转换，因为myInt和int是两种数据类型
fmt.Println("k1,k2",k1, k2)
*/
