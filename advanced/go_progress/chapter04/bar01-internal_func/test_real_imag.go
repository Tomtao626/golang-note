package main

import "fmt"

func main() {

	var com complex64

	com = 12.11 + 14i

	fmt.Println(com)
	fmt.Println(real(com))
	fmt.Println(imag(com))
}
