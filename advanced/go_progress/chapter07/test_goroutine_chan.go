package main

import (
	"fmt"
	"sync"
)

// 猫主子和铲屎官日常

var Wg = sync.WaitGroup{}

// 铲屎官负责投食
func Man(ch chan string,foot string)  {
	ch <- foot
	fmt.Printf("投食了:%s\n",foot)
	Wg.Done()
}

// 猫主子负责吃食
func Cat(ch chan string)  {
	foot,ok := <- ch

	fmt.Println(ok)
	if ok {
		fmt.Printf("吃了：%s\n",foot)
	}else {
		fmt.Println("铲屎官该投食了")
	}
	Wg.Done()

}

func main() {


	ch := make(chan string)
	defer close(ch)
	foot := "猫罐头"

	for i:=0;i<10;i++ {
		Wg.Add(2)
		go Man(ch,foot)
		go Cat(ch)
	}


	Wg.Wait()
}
