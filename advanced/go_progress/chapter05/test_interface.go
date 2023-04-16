package main

import "fmt"

type MyInterface interface {
	Say() // 规范好函数名称
	Run()
}

func Em(m MyInterface)  {
	m.Say()
	//m.Run()
}

type Person1 struct {
}

func (p Person1)Say()  { // 具体实现

	fmt.Println("这是say方法")
}

func (p Person1)Run()  {
	fmt.Println("这是run方法")

}


func main() {


	//map_data := map[string]interface{}{"name":"hallne","age":18}

	person := Person1{}
	Em(person)
}

