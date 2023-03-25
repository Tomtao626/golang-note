package main

import "fmt"

// 求函数值
// 已知 f(1) = 3, f(n) = 2*f(n-1)+1, 运用递归的思想，求出f(n)的值

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func main() {
	fmt.Println(f(1))
	fmt.Println(f(2))
}
