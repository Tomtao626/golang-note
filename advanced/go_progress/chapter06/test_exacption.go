package main

import (
	"fmt"
	"strconv"
)

func Arr(slice []int)  {
	fmt.Println(slice[3]) // 长度小于4会报错
}


func main() {
	//slice := []int{1,2,3}
	//Arr(slice)


	//var num *int
	//num = new(int)
	//*num = 3
	//fmt.Println(num)
	//fmt.Println(*num)


	i,err := strconv.ParseInt("aa",10,64)
	fmt.Println(i)
	//fmt.Println(err)
	if err != nil {
		fmt.Println("转换失败")
		panic("转换失败")
	}


	//
	fmt.Println("正常运行")
	//
	//
	//panic("不让正常运行")


}

