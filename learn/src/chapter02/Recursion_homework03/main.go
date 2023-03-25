package main

import "fmt"

// 猴子吃桃子问题
// 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个，以后每天都吃前一天剩下的一半多一个。
// 当吃到第十天时，想再吃时（还没吃），发现只有一个桃子了，问：最初共有多少个桃子？
/*
思路分析：
	1.第十天只有一个桃子了
	2.第九天还剩几个桃子？第九天桃子数 = (第十天桃子数+1)*2
	3.前一天桃子数等于(第二天桃子数+1)的2倍
	4. 即 peach(n) = (peach(n+1)+1)*2
*/

// 递归写法

func peach(n int) int {
	if n >= 10 || n < 1 {
		fmt.Println("输入天数不对，请重新输入！")
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	peachNum := peach(1)
	fmt.Println("peach_num is :", peachNum)
}

// 屌丝写法
/*
func main(){
	x , y ,day := 0, 1, 9
	for day > 0 {
		x = (y + 1) * 2
		y = x
		day--
	}
	fmt.Println("peachNum is :", y)
}
*/
