package main

import (
	"strconv"
	"fmt"
)

func main() {

	ret,err := strconv.ParseInt("xx",10,64)

	fmt.Println(ret)
	panic(err)

	recover()
	//fmt.Println(err)
}
