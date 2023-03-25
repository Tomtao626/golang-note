package main

import "fmt"

// 用 Golang 实现，将四个数进行排列组合。
// 有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
/*
思路分析：
	可填在百位、十位、个位的数字都是 1、2、3、4。组成所有的排列后再去掉不满足条件的排列。
	暴力法 三层遍历
*/
func allArray() int {
	totalNum := 0
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				// 判断i，j，k三个都不相同
				if i != j && j != k && i != k {
					totalNum++
					fmt.Println("第", totalNum, "方案", "i =", i, "j =", j, "k =", k)
				}
			}
		}
	}
	fmt.Println("the totalnum is :", totalNum)
	return totalNum
}

func main() {
	allArray()
}
