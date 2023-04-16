package main

import "fmt"

type Person struct {
}

type Base interface {
}

// 类型别名
type Myint int

// 定义函数类型
type Say func(name string) string

func Hello(name string)string  {
	fmt.Println("==========")
	ret := fmt.Sprintf("hello %s",name)
	return ret
}

func (s Say)Hi(n string)string  {

	ret := fmt.Sprintf("say:%s",s(n))
	return ret

}

func main() {

	var i Myint
	i = 1
	fmt.Println(i)
	fmt.Printf("%T\n",i)

	ret := Say(Hello).Hi("hallen")
	fmt.Println(ret)


}
