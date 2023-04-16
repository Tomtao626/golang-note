package main

import (
	"sync"
	"fmt"
	"time"
)


var wg = sync.WaitGroup{

}

func Hi2()  {
	fmt.Println("hi2执行了")
	wg.Done()
}

func Hi1()  {
	fmt.Println("hi1执行了")
	wg.Done()

}


func main() {


	start_time := time.Now()
	fmt.Println(start_time)

	wg.Add(2)

	go Hi1()
	go Hi2()

	wg.Wait()

	end_time := time.Now()
	fmt.Println(end_time)
	fmt.Println(end_time.Sub(start_time))


}
