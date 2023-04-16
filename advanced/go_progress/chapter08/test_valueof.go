package main

import (
	"fmt"
	"reflect"
)

func SetValue(i int)  {
	fmt.Println(i)

	v := reflect.ValueOf(&i)

	v.Elem().SetInt(4)
	fmt.Println(i)

}
func Get(param interface{})  {

	v := reflect.ValueOf(param)

	fmt.Println(v)
	fmt.Println(v.Kind())
	fmt.Println(v.String())
	fmt.Println(v.Type())
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		fmt.Println("这是int类型")
	case reflect.String:
		fmt.Println("这是string类型")



	}

}
func main() {

	n  := 23

	//Get(n)
	SetValue(n)
	//fmt.Println(n)
	//fmt.Println(reflect.ValueOf(n))
}
