package main

import (
	"sync"
	"fmt"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex
	mu.Lock()
	mu.count++
	mu.Unlock()
	var mu1 = mu
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}

//func main() {
//
//	ch := make(chan int)
//
//	ch <- 3
//}
