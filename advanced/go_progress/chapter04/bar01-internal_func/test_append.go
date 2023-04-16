package main

import "fmt"

func main() {

	// 数组不能追加
	arr := [5]int{1,2,3}
	arr[4] = 5

	//arr = append(arr,6)
	fmt.Println(arr)
	// 切片

	slice := []int{1,2,3}

	slice1 := []int{}
	//slice[2] = 4
	// 使用新的切片变量接收
	//slice1 = append(slice,5 )

	// 使用原来的切片接收
	slice = append(slice,5 )
	//fmt.Println(slice)
	fmt.Println(slice1)



}
