package main

import (
	"sync"
	"fmt"
	"time"
)

var sum1 int

var rwmutex sync.RWMutex

func Write()  {
	rwmutex.Lock()
	sum1 ++
	fmt.Printf("写入数据正常，加完后的值是：%d\n",sum1)
	rwmutex.Unlock()
}

func Read()  {
	rwmutex.RLock()
	fmt.Printf("读数据正常，读到的值是：%d\n",sum1)
	rwmutex.RUnlock()
}


func main() {

	//swg1.Add(2)
	for i:=0;i<10;i++{
		go Write()
	}
	for j:=0;j<10;j++ {
		go Read()
	}

	time.Sleep(time.Second * 20)
	//swg1.Wait()
}
