package main

import (
	"fmt"
	"learn_go/chapter02/func_demo01/utils"
)

func main() {
	var n1 float64 = 1.8
	var n2 float64 = 2.9
	var operator byte = '+'
	results := utils.Cal(n1, n2, operator)
	fmt.Println(results)
}
