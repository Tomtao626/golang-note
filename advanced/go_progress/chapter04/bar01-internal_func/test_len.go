package main

import "fmt"

func main() {


	// string
	str := "hallen"
	fmt.Println(len(str))

	// array
	arr := [5]int{}
	fmt.Println(len(arr))

	// slice
	slice := []int{1,2,3,4}
	fmt.Println(len(slice))

	// map
	map_data := map[string]int{
		"name":111,
		"age":222,
		"addr":444,
	}

	fmt.Println(len(map_data))

	// chan
	var ch = make(chan int) // 无缓存的chan
	var ch1 = make(chan int,10) // 有缓存的chan
	ch1 <- 3
	ch1 <- 2
	fmt.Println(len(ch))
	fmt.Println(len(ch1))
	fmt.Println(cap(ch1))


}
