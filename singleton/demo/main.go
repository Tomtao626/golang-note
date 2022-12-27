package main

import (
	"fmt"
	"sync"
)

type Woman struct {
	name string
}

var (
	once sync.Once
	ting *Woman
)

func getTing() *Woman {
	once.Do(func() {
		ting = new(Woman)
		ting.name = "tingting"
		fmt.Println("newtingting")
	})
	fmt.Println("gettingting")
	return ting
}

func main() {
	for i := 0; i < 3; i++ {
		_ = getTing()
	}
}
