package main

import "fmt"

func main() {


	//var slice []int = make([]int,3)
	var slice []int = make([]int,3,10) // 长度为10 的数组

	slice[0] = 1
	//slice[10] = 111
	fmt.Println(slice)

	// panic: runtime error: index out of range [0] with length 0
	// panic: runtime error: index out of range [3] with length 3
	//s := []int{1,2,3}

	//fmt.Println(s[3])


	// map

	// panic: assignment to entry in nil map
	var map_data map[string]interface{} = make(map[string]interface{})

	map_data["name"] = "hallen"
	fmt.Println(map_data)

	// chan

	// fatal error: all goroutines are asleep - deadlock!
	var ch chan int = make(chan int,1)

	ch <- 3
	fmt.Println(<- ch)

	ch <- 4
	fmt.Println(<- ch)
	ch <- 5


}
