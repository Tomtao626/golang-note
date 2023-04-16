package main

func main() {

	var ch = make(chan int) // 无缓存的chan

	close(ch)

	ch <- 3
}
