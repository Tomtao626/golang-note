package main

import "fmt"

func main() {

	name := "hallen"
	var new_name *string = &name

	fmt.Println(&new_name)

	*new_name = ""
	fmt.Println(&new_name)

	new_name = nil
	fmt.Println(&new_name)


}
