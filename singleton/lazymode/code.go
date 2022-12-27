package main

import (
	"sync"
)

type Singleton struct{}

var instance *Singleton
var once sync.Once

// GetInstance 懒汉模式单例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	GetInstance()
}
