package main

import "fmt"

func Slice(slice []int)  {
	slice[1] = 111
}

func main() {
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)
	Slice(slice)
	fmt.Println(slice)

	var slice2 []int = make([]int,10)
	slice2[0] = 11
	fmt.Println(slice2)
}
