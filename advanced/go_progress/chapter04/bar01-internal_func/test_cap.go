package main

import "fmt"

func main() {


	var slice []int = make([]int,2,4) // 多余两个空间大小
	var slice1 []int = make([]int,2)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
}
