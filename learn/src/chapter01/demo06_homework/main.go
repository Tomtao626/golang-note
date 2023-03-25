/*
	1.还有97天放假，问：折算成是多少个星期零多少天
	2.定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)， 请求出华氏温度对应的摄氏温度
*/
package main

import "fmt"

func main() {
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("还有%d个星期零%d天\n", week, day)

	var huashi float32 = 178.34
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度%v 对应的摄氏温度为：%v", huashi, sheshi)
}
