package main

import "fmt"

func getSum(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	var x int = 7
	var y int = 8
	sum := getSum(x, y)
	fmt.Println("sum is: ", sum)
}
