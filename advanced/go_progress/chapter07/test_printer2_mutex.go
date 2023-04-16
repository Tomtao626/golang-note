package main

import (
	"sync"
	"fmt"
)

var swg = sync.WaitGroup{}
var sum int

var mutex1 sync.Mutex

func Sum()  {

	mutex1.Lock()
	for i:=0;i<10;i++{
		fmt.Println(sum)
		sum ++
	}
	mutex1.Unlock()
	swg.Done()

}
func main() {

	swg.Add(2)
	go Sum()
	go Sum()

	swg.Wait()
}
