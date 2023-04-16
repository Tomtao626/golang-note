package main

import (
	"fmt"
	"time"
	"sync"
)

var mwg = sync.WaitGroup{}

var mutex sync.Mutex


func Printer(str string)  {

	mutex.Lock()
	for i:=0;i<len(str);i++ {
		//fmt.Println(str[i])
		fmt.Printf("%c",str[i])
		time.Sleep(time.Second)
	}
	mutex.Unlock()
}


func Person1()  {
	str := "hello"
	Printer(str)
	mwg.Done()
}

func Person2()  {
	str := "zhiliao"
	Printer(str)
	mwg.Done()

}
func main() {

	mwg.Add(2)
	go Person1()
	go Person2()

	mwg.Wait()


}
