package main

import "fmt"

func Arr(arr [5]int)  {

	arr[2] = 111
	fmt.Println(arr)

}
func main() {

	arr := [5]int{1,2,3}
	fmt.Println(arr)
	Arr(arr)
	fmt.Println(arr)

}
