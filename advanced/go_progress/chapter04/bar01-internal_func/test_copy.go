package main

import (
	"fmt"
)

func main() {

	// dst长度大于src长度
	dst := []int{5,6,7}
	src := []int{4,3}
	count := copy(dst,src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(count)

	// dst长度小于src长度
	dst1 := []int{5}
	src1 := []int{4,3}
	count1 := copy(dst1,src1)
	fmt.Println(dst1)
	fmt.Println(src1)
	fmt.Println(count1)

	// dst 和 src的长度相等
	dst2 := []int{5,6,7}
	src2 := []int{4,3,2}

	copy(dst2,src2)
	fmt.Println(dst2)
}
