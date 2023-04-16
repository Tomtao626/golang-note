package main

import (
	"fmt"
	"errors"
)

// 错误的原因只有一个
func Add(name string) error {

	if name != "halle" {

		return errors.New("name is not hallen")
	}
	return nil

}

type Person2 struct {
	Name string
}
// 不会发生错误
func (p *Person2)GetName(str string)  error{
	p.Name = str
	return nil
}

func main() {

	defer recover()
	arr := []int{1,2,3}

	fmt.Println(arr[3])

}
