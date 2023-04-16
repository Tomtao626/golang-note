package main

import "fmt"

func main() {

	// chan进行数据通信
	//var ch chan int


	// 读写通道
	//ch := make(chan int)

	// 只读通道
	//ch1 := make(<-chan int)

	// 只写通道
	//ch2 := make(chan <- int)

	// 写入数据
	//go func() {
	//	ch <- 3
	//}()

	// 读数据
	//fmt.Println(<- ch)


	// 有缓存的通道

	ch3 := make(chan int,2)
	defer close(ch3)
	defer func() {
		recover()
	}()
	ch3 <- 3
	ch3 <- 4
	//fmt.Println(<- ch3)
	//ch3 <- 5
	//fmt.Println(<- ch3)
	//fmt.Println(<- ch3)
	//fmt.Println(<- ch3)
	//fmt.Println(<- ch3)
	//fmt.Println(len(ch3))
	//for i:=0;i<len(ch3);i++ {
	//	fmt.Println(<- ch3)
	//	//fmt.Println(<- ch3)
	//}

	for val := range ch3{
		fmt.Println(val)
	}



}
