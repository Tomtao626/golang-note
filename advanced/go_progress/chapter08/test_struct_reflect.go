package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Id   int
	Name string `json:"name"`
	Age  int
	Addr string
}

func (p *Person) Add() {
	fmt.Println("这是Add方法")
}
func (p *Person) Get() {
	fmt.Println("这是Get方法")

}

func main() {

	person := Person{Id: 1, Name: "hallen", Age: 18, Addr: "xxx"}

	//t := reflect.TypeOf(person)
	//v := reflect.ValueOf(person)

	//fmt.Println(t)
	//fmt.Println(v)

	// 字段信息
	//fmt.Println(t.NumField())
	//fmt.Println(t.Field(0))
	//fmt.Println(t.Field(1))
	//fmt.Println(t.Field(4))
	//fmt.Println(t.FieldByName("Name"))
	//for i:=0;i<t.NumField();i++ {
	//	fmt.Println(t.Field(i))
	//}

	// method
	t := reflect.TypeOf(&person)
	fmt.Println(t.NumMethod())
	//fmt.Println(t.Method(0))
	//fmt.Println(t.Method(1))
	//fmt.Println(t.Method(2))
	fmt.Println(t.MethodByName("Add"))
	t2, _ := t.MethodByName("Add")
	t2.Func.Call([]reflect.Value{reflect.ValueOf(&person)})
	//for i:=0;i<t.NumMethod();i++ {
	//	fmt.Println(t.Method(i))
	//}
	//
	//person.Get()

	// 通过字符串的方法名称调用该方法，比如："Get"

	v := reflect.ValueOf(&person)
	v2 := v.MethodByName("Get1")

	// 调用
	v2.Call([]reflect.Value{})

}
