package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	nums1 = append(nums1, nums2...)
	return nums1
}

// fbn 斐波那契数列
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

// BubbleSort 冒泡排序
func BubbleSort(arr *[5]int) {
	//fmt.Println("排序前arr=", (*arr))
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	//fmt.Println("排序后arr=", (*arr))
}

// SeqSearch 顺序查找
func SeqSearch(arr *[6]int, findVal int) {
	// 遍历数组
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == findVal {
			fmt.Printf("找到了，下标为%v\n", i)
			break
		} else if i == len(*arr)-1 {
			fmt.Println("找不到")
		}
	}
}

// SeqSearchPro 顺序查找高级写法
func SeqSearchPro(arr *[6]int, findVal int) {
	// 遍历数组
	for i, v := range arr {
		if v == findVal {
			fmt.Printf("找到了，下标为%v\n", i)
			break
		} else if i == len(arr)-1 {
			fmt.Println("找不到")
		}
	}
}

// BinaryFind 二分查找
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		// 要查找的数，应该在leftIndex~middle-1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 要查找的数，应该在middle+1~rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Printf("找到了，下标为%v\n", middle)
	}
}

// SelectSort 选择排序
func SelectSort(arr *[5]int) {
	// 完成第一轮排序
	// 假设arr[0]是最小值
	for j := 0; j < len(*arr)-1; j++ {
		min := (*arr)[j]
		minIndex := j
		// 遍历后面的1~len(arr)-1比较
		for i := j + 1; i < len(*arr); i++ {
			if min > (*arr)[i] {
				// 修改min minIndex
				min = (*arr)[i]
				minIndex = i
			}
		}
		// 交换
		if minIndex != j {
			(*arr)[j], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[j]
		}
		fmt.Printf("第%d次%v\n", j+1, *arr)
	}
}

// InsertSort 插入排序
func InsertSort(arr *[5]int) {
	// 完成第一轮排序
	// 假设arr[0]是最小值
	for j := 1; j < len(*arr); j++ {
		insertVal := (*arr)[j]
		insertIndex := j - 1
		// 从大到小
		for insertIndex >= 0 && (*arr)[insertIndex] < insertVal {
			(*arr)[insertIndex+1] = (*arr)[insertIndex]
			insertIndex--
		}
		// 插入
		if insertIndex+1 != j {
			(*arr)[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次%v\n", j, *arr)
	}
}

// QuickSort 快速排序
func QuickSort(left int, right int, arr *[5]int) {
	l := left
	r := right
	// pivot 中轴
	pivot := (*arr)[(left+right)/2]
	temp := 0
	// for循环的目标是将比pivot小的数放到左边
	// 比pivot大的数放到右边
	for l < r {
		// 从pivot的左边找到大于等于pivot的值
		for (*arr)[l] < pivot {
			l++
		}
		// 从pivot的右边找到小于等于pivot的值
		for (*arr)[r] > pivot {
			r--
		}
		// l >= r说明本次分解任务完成，break
		if l >= r {
			break
		}
		// 交换
		temp = (*arr)[l]
		(*arr)[l] = (*arr)[r]
		(*arr)[r] = temp
		// 优化
		if (*arr)[l] == pivot {
			r--
		}
		if (*arr)[r] == pivot {
			l++
		}
	}
	// 如果l == r，必须l++，r--，否则会出现栈溢出
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	//s1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//s2 := []int{2, 5, 6}
	//n := 3
	//s1 = merge(s1, m, s2, n)
	//sort.Ints(s1)
	//fmt.Println(s1)

	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println(arr)
}
