package main

import "fmt"

func main() {
	// 定义局部变量
	var a int = 10
	// 循环♻️
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			// 跳过迭代
			goto LOOP
		}
		fmt.Printf("a is %d\n", a)
		a++
	}
}

/*
out:
	a is 10
	a is 11
	a is 12
	a is 13
	a is 14
	a is 16
	a is 17
	a is 18
	a is 19
*/
