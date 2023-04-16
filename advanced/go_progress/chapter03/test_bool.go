package main

import "fmt"

func IsTrue(b bool)  {
	if !b {
		fmt.Println("这是false")
	}else {
		fmt.Println("这是true")
	}
}
func main() {

	b := true
	IsTrue(b)

}
