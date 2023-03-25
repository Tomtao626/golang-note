package main

import "fmt"

func main() {
	a := 9
	b := 10
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	flag := a == b
	fmt.Println("flag:", flag)
}
