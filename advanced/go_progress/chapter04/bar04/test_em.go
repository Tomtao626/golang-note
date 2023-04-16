package main

import (
	"fmt"
	"golang-note/advanced/go_progress/chapter04/bar04/base"
)

type Student struct {
	base.MyBase
	ClaNum int
	Course int
	StuNum int
}

func (s *Student) Study() {
	fmt.Println("我会学习")

}

func main() {

	student := Student{ClaNum: 2, Course: 98, StuNum: 111, MyBase: base.MyBase{Id: 222, Name: "张三"}}

	//student.Id = 222
	//student.Name = "张三"

	fmt.Println(student)
	student.Study()
	student.Run()

}
