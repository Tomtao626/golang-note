package main

import (
	"fmt"
	"unsafe"
)

type Part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

func main() {
	var a bool
	var b int32
	var c int8
	var d int64
	var e byte
	fmt.Printf("bool size: %d\n", unsafe.Sizeof(a))
	fmt.Printf("int32 size: %d\n", unsafe.Sizeof(b))
	fmt.Printf("int8 size: %d\n", unsafe.Sizeof(c))
	fmt.Printf("int64 size: %d\n", unsafe.Sizeof(d))
	fmt.Printf("byte size: %d\n", unsafe.Sizeof(e))
	fmt.Printf("string size: %d\n", unsafe.Sizeof("EDDYCJY"))

	part := Part1{}

	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part), unsafe.Alignof(part))
}