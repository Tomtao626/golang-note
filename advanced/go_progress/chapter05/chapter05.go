package main

import (
	"fmt"
)

type A interface {
	B() // 只是规定好需要实现的方法名称
}

type C interface {
	B()
	D()
}


func GetA(a A)  {
	a.B()
}

type Student struct {}

func (s Student) B()  {


}

type Teacher struct {
	Name string
}

func F(v interface{})  {
	man := v.(Teacher)
	//t := Teacher{}

	switch v {
	case Teacher{}:
		man.Name = "hallen"

	}

}
func main() {

	type Myint int
	var mi Myint

	var i int = 1

	//mi = i

	mi = Myint(i)
	//
	//var in interface{}
	//
	//mi = in.(Myint)
	fmt.Println(mi)


	a := interface{}(nil)
	b := (*int)(nil)

	_,c := a.(interface{})

	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Println(a == b)
	fmt.Println(b == nil)  // keng
	fmt.Println(c)


	var i1 interface{}
	var i2 interface{}

	if i1 == i2 {
		fmt.Println("相等")
	}

	// todo 指针类型转换
	//var num int = 3
	//var aa *int = &num
	//
	//var bb *int64
	//
	//bb = (*int64)(aa)

}