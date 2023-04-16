package main

import "fmt"

func main() {


	map_data := map[string]interface{}{
		"name":"hallen",
		"age":18,
	}


	fmt.Println(map_data)
	delete(map_data,"name")
	fmt.Println(map_data)
}
