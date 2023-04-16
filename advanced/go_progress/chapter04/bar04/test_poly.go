package main

import "fmt"

type Animal interface{
	Say()
}

func AnimailSay(a Animal)  {
	a.Say()
}


type Cat struct {
	Name string
	Age int
}

func (c *Cat)Say()  {
	fmt.Printf("%s,年龄:%d,喵喵的叫",c.Name,c.Age)
}

type Dog struct {
	Name string
	Age int
}

func (d *Dog)Say()  {
	fmt.Printf("%s,年龄:%d,汪汪的叫",d.Name,d.Age)
}



func main() {

	dog := Dog{Name:"小狗",Age:2}
	cat := Cat{Name:"小花",Age:1}


	AnimailSay(&dog)
	AnimailSay(&cat)


}
