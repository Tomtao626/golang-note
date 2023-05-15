package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	nums1 = append(nums1, nums2...)
	return nums1
}

func fbn(n int) []uint64 {
	// 声明一个切片 切片大小为n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	// for循环来存放斐波拉契数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	//s1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//s2 := []int{2, 5, 6}
	//n := 3
	//s1 = merge(s1, m, s2, n)
	//sort.Ints(s1)
	//fmt.Println(s1)
	fbnSlice := fbn(10)
	fmt.Println(fbnSlice)
}
