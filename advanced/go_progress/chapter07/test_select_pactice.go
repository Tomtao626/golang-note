package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int,5)

	for i:=0;i<10;i++ {
		select {
		case ch <- 3:
			fmt.Printf("写入操作第%d次\n",i+1)
			// 超时判断
		case <- time.After(time.Second * 2):
			fmt.Println("写入超时")
		}
	}

}
