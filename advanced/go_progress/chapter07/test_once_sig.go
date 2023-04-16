package main

import (
	"fmt"
	"sync"
)


var sonce sync.Once



type Person struct {
	Name string
}

var person *Person


func NewPerson(name string)*Person  {

	sonce.Do(func() {
		person = new(Person)
		person.Name = name

	})
	//person := Person{Name:name}
	return person
}

func main() {

	//p1 := Person{Name:"hallen1"}
	//p2 := Person{Name:"hallen2"}



	p1 := NewPerson("hallen1")
	p2 := NewPerson("hallen2")
	fmt.Printf("%p\n",p1)
	fmt.Printf("%p\n",p2)

}
