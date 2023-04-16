package main

import (
	"fmt"
	"strconv"
)



func main() {

	// 使用err代替panic

	num,err := strconv.ParseInt("aa",10,64)
	fmt.Println(num)
	fmt.Println(err)

	defer func() {
		err := recover()
		if err !=nil {
			fmt.Println(err)
			return
		}
	}()
	var i *int
	//i = new(int)
	*i = 3


	fmt.Println("正常运行")


}
