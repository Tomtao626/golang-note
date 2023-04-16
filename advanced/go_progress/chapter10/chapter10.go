package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	b bool
	a int
	c int8
	d int32
	e byte
}


func main() {

	var a int
	var b bool
	var c int8
	var d int32
	var e byte


	fmt.Printf("a int size:%v\n",unsafe.Sizeof(a))
	fmt.Printf("a int alignof:%v\n",unsafe.Alignof(a))

	fmt.Printf("b bool size:%v\n",unsafe.Sizeof(b))
	fmt.Printf("b bool alignof:%v\n",unsafe.Alignof(b))

	fmt.Printf("c int8 size:%v\n",unsafe.Sizeof(c))
	fmt.Printf("d int32 size:%v\n",unsafe.Sizeof(d))
	fmt.Printf("e byte size:%v\n",unsafe.Sizeof(e))

	user := User{}
	fmt.Printf("user struct size:%v\n",unsafe.Sizeof(user))
}
