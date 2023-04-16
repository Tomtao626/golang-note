package main

import (
	//"fmt"
	//"reflect"
	"fmt"
	"reflect"
	"strconv"
)

// 在i的基础+1
func Count(i interface{}) int{

	t := reflect.TypeOf(i)
	fmt.Println(t)
	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.String())
	//n := i.(int)
	//n++
	//return n

	kind := t.Kind()

	fmt.Println(kind)

	switch kind {
	case reflect.Int:
		ret := i.(int)
		ret++
		return ret
	case reflect.String:
		ret,_ := strconv.ParseInt(i.(string),10,64)
		ret ++
		return int(ret)
	default:
		return 0
	}


}
func main() {

	//var I int = 1
	//var name string
	//
	//fmt.Println(reflect.TypeOf(I))
	//fmt.Println(reflect.TypeOf(name))

	i := "2"
	n := Count(i)
	fmt.Println(n)
}



