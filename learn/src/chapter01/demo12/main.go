package main

import "fmt"

func main() {
	//层数
	var totallevel = 20

	for i := 1; i <= totallevel; i++ {
		// 在打印*前 先打印空格
		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}
		// j 表示每层打印多少
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || j == totallevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
