package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello hallen"

	fmt.Println(str)

	// 转大写
	fmt.Println(strings.ToUpper(str))

	// 转小写

	ret := strings.ToUpper(str)
	fmt.Println(strings.ToLower(ret))

	// fields：根据空格
	fmt.Println(strings.Fields(str))

	// Join
	ret2 := strings.Fields(str)
	fmt.Println(strings.Join(ret2,"-"))

	// count
	fmt.Println(strings.Count(str,"h"))
	fmt.Println(strings.Count(str,"l"))


	// replace
	old := "h"
	new_substr := "w"

	fmt.Println(strings.Replace(str,old,new_substr,1))
	fmt.Println(strings.Replace(str,old,new_substr,-1))

	// replaceAll
	fmt.Println(strings.ReplaceAll(str,old,new_substr))

	//HasPrefix
	fmt.Println(strings.HasPrefix(str,"h"))
	fmt.Println(strings.HasPrefix(str,"w"))

	// HasSuffix
	fmt.Println(strings.HasSuffix(str,"n"))
	fmt.Println(strings.HasSuffix(str,"a"))

	var arr [5]int
	fmt.Println(arr)

	slice := []int{}

	fmt.Println(slice)

}
