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

	fmt.Printf("a int size:%v\n", unsafe.Sizeof(a))
	fmt.Printf("a int alignof:%v\n", unsafe.Alignof(a))

	fmt.Printf("b bool size:%v\n", unsafe.Sizeof(b))
	fmt.Printf("b bool alignof:%v\n", unsafe.Alignof(b))

	fmt.Printf("c int8 size:%v\n", unsafe.Sizeof(c))
	fmt.Printf("d int32 size:%v\n", unsafe.Sizeof(d))
	fmt.Printf("e byte size:%v\n", unsafe.Sizeof(e))

	user := User{}
	fmt.Printf("user struct size:%v\n", unsafe.Sizeof(user))
}

/*
a 对齐值 8, 8字节对齐, 64位系统, 8字节, 32位系统, 4字节, 偏移量为 0
b 对齐值 1, 1字节对齐, 64位系统, 1字节, 32位系统, 1字节, 偏移量为 8 从9开始
c 对齐值 1, 1字节对齐, 64位系统, 1字节, 32位系统, 1字节, 偏移量为 9 从10开始
d 对齐值 4, 4字节对齐, 64位系统, 4字节, 32位系统, 4字节, 偏移量为 12 从13开始
e 对齐值 1, 1字节对齐, 64位系统, 1字节, 32位系统, 1字节, 偏移量为 14 从17开始

aaaa|aaaa|bc..|dddd|e  17
申请内存为 24  因为16字节不够 24字节对齐

user struct size:16


*/
