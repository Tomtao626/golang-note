package main

import (
	"fmt"
	"math"
)

// 用 Golang 实现，计算完全平方数。
// 一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？
/*
思路分析：
	在 10 万以内判断，先将该数加上 100 后再开方，再将该数加上 268 后再开方，如果开方后的结果满足如下条件，即是结果。
*/

func main() {
	i := 0
	for {
		x := int(math.Sqrt(float64(i + 100)))
		y := int(math.Sqrt(float64(i + 268)))
		if x*x == i+100 && y*y == i+268 {
			fmt.Printf("the num is %d\n", i)
			break
		}
		i++
	}
}
