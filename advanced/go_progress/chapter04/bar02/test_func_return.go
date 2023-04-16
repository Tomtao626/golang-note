package main

import "fmt"

// 返回一个值：第一种方式
func Add(x,y int) int {

	ret := x + y
	fmt.Println(ret)
	return ret
}


// 返回一个值：第二种方式
func Add1(x,y int) (ret int) {
	ret = x + y
	fmt.Println(ret)
	return ret
}

// 返回多个值：第一种方式
func MyMath(x,y int) (int,int,int,int) {

	sum := x + y
	sub := x - y
	mult := x * y
	divi := x / y
	return sum,sub,mult,divi

}

// 返回多个值：第二种方式
func MyMath1(x,y int) (sum ,sub,mult,divi int) {

	sum = x + y
	sub = x - y
	mult = x * y
	divi = x / y
	return

}


func main() {

	//sum := Add(2,3)
	//fmt.Println(sum)
	//
	//Add1(4,5)

	sum ,sub,mult,divi := MyMath1(4,5)
	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mult)
	fmt.Println(divi)



}
