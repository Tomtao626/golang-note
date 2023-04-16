package main

import "fmt"

func init() {

}


func add(args ...int) int {
	sum :=0

	for _,arg := range args {
		sum += arg

	}
	return sum

}

func main() {


	//map_data := map[string]interface{}{
	//	"name":"hallen",
	//}

	//cap(map_data)  // 不能对map类型cap计算

	var ch chan int

	num := cap(ch)
	fmt.Println(num)


	slice := make([]int,10)  // 指针


	slice = append(slice, 1,2,3)
	fmt.Println(slice)


	// 什么是闭包

	// 特点

	slice1 := []int{1,2,3,4,5}
	slice2 := []int{5,6,7,8,9,10}

	slice1 = append(slice1, slice2...)

}

func Sum(num ...[]int)  {

}
