package main

//go函数本身也是一种数据类型，可以赋值给一个变量，该变量就是一个函数类型的变量
import "fmt"

func getTest(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	a := getTest
	fmt.Printf("a的类型%T, getTest的类型是%T\n", a, getTest)

	res := a(9, 10) // 等价于 res := getTest(9, 10)
	fmt.Println("res=", res)
}
