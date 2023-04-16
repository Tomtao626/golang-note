package main

import (
	"reflect"
	"fmt"
)

type Person1 struct {

	Id int `json:"id" hallen:"user_id" `
	Name string `json:"name" hallen:"user_name"`
}

func main() {

	person1 := Person1{}
	t := reflect.TypeOf(&person1)

	for i:=0;i<t.Elem().NumField();i++ {
		field := t.Elem().Field(i)
		j := field.Tag.Get("json")
		h := field.Tag.Get("hallen")
		fmt.Println(field)
		fmt.Println(j)
		fmt.Println(h)

	}





}
