package main

import "fmt"

func main() {

	ch := make(chan int,2)

	go func() {
		ch <- 3
	}()

	//ch := make(chan int)

	//a := 1
	//b := 1

	select {
		//case a == b:
		//	fmt.Println("判断相等操作")
		case num := <- ch:
			fmt.Println(num)
			fmt.Println("读操作没问题")
			fmt.Println("读操作退出")
		case ch <- 4:
			fmt.Println("写操作没问题")
			break
			fmt.Println("写操作退出")
		default:
			fmt.Println("default操作")


	}


}
