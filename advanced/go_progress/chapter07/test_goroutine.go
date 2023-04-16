package main

import (
	"fmt"
	"time"
)

func Hello2()  {
	fmt.Println("hello2函数执行了")
}
func Hello()  { // 任务
	fmt.Println("hello函数执行了")
}


func main() {
	// main函数是个协程，而且是主协程



	fmt.Println("main函数执行了") // 0.01

	// 协程创建时需要时间的
	start_time := time.Now()
	fmt.Println(start_time)
	go Hello()  // 0.02
	go Hello2()
	end_time := time.Now()
	fmt.Println(end_time)
	sub_time := end_time.Sub(start_time)
	fmt.Println(sub_time)
	//fmt.Println(sub_time.Microseconds())
	//fmt.Println(sub_time.Milliseconds())

	time.Sleep(time.Second * 1)

}

