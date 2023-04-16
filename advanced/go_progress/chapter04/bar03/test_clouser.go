package main

import "fmt"

func Add1() func() {
	x := 0 // 局部变量
	f := func(){
		x ++
		fmt.Println(x)
	}

	return f


}

func main() {
	f := Add1()

	f()
	f()
	f()

}
