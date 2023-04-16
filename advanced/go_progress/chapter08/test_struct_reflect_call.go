package main

import (
	"fmt"
	"reflect"
)

type MyMath struct {
	X int
	Y int
}

// 加法
func (m *MyMath) Add() int {
	return m.X + m.Y

}

// 减法
func (m *MyMath) Sub() int {
	return m.X - m.Y

}

// 乘法
func (m *MyMath) Muli() int {
	return m.X * m.Y

}

// 除法
func (m *MyMath) Divi() int {
	return m.X / m.Y

}

func main() {



	var operation string
	var x int
	var y int

	// 让用户输入
	fmt.Println("请输入第一个值")
	fmt.Scanln(&x)
	fmt.Println("请输入第二个值")
	fmt.Scanln(&y)
	//fmt.Println("请输入要运算的方法，+：Add，-：Sub，*：Muli，/：Divi")
	fmt.Println("请输入要运算的符号，+，-，*，/")
	fmt.Scanln(&operation)
	mymath := MyMath{X:x,Y:y}

	map_data := map[string]string{
		"+":"Add",
		"-":"Sub",
		"*":"Muli",
		"/":"Divi",
	}

	name := map_data[operation]



	v := reflect.ValueOf(&mymath)

	v2 := v.MethodByName(name)

	v3 := v2.Call([]reflect.Value{})
	for _,v4 := range v3 {
		fmt.Println(v4.Int())
	}

}
