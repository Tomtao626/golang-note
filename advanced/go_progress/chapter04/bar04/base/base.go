package base

import "fmt"

type MyBase struct {
	Id int
	Name string
}

func (m *MyBase) Run()  {

	fmt.Println("跑起来了")

}


