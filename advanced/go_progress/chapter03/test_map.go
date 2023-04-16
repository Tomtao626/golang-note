package main

import "fmt"

func Map(data map[string]string)  {

	data["name"] = "zhiliao"
}

func main() {
	data := map[string]string{
		"name":"hallen",
		"age":"18",
	}

	fmt.Println(data)
	Map(data)
	fmt.Println(data)

	var data2 map[string]string = make(map[string]string)

	data2["age"] = "111"
	fmt.Println(data2)

}
