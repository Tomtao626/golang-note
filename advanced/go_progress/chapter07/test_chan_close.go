package main

import (
	"time"
	"fmt"
)

func main() {

	var ch chan int = make(chan int,2)

	//ch <- 3
	close(ch)
	//ch <- 3

	//go func() {
	//	ch <- 3
	//}()

	num,ok := <- ch

	if ok {
		fmt.Println(num)
	}
	//go func() {
	//	fmt.Println(<- ch)
	//}()

	time.Sleep(time.Second)

}
