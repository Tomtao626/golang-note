package main

import "fmt"

// 9*9
func main() {
	Print9X9()
}

// for循环-实现99乘法口诀表
func Print9X9() {
	var num = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}

// goto-实现99乘法口诀表
func goto9X9() {

}
