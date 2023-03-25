package main

import "fmt"

// 斐波拉契数列 1 1 2 3 5 8 13
// 当 n == 1 || n == 2 时，返回 1
// 当 n >= 2 时，返回前面两个数的和 f(n-1)+f(n-2)

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	fmt.Println(fbn(3))
	fmt.Println(fbn(4))
	fmt.Println(fbn(5))
	fmt.Println(fbn(6))
	fmt.Println(fbn(7))
}
