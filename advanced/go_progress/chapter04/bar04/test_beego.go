package main

import (
	"fmt"
	"golang-note/advanced/go_progress/chapter04/bar04/base"
)

type User struct {
	base.Base
	Id       int
	UserName string
	Age      int
	Avater   string
}

func (u *User) Get() {
	u.BaseRender()
	u.Age = 18
	u.Avater = "xxx.jpg"

	fmt.Println(u)

}
func main() {

	user := User{}
	user.Get()

}
