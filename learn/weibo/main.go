package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	urlFmt    = "https://s.weibo.com/ajax/jsonp/gettopsug?uid=&ref=PC_topsug&url=https%%3A%%2F%%2Fs.weibo.com%%2Ftop%%2Fsummary&source=&aid=01AanI52aVi_Fj0FjPCIWggN34vvVC0NMD9LevrsZVf57HYsM.&_rnd=%d"
	threadNum = 1
)

// 1. 通过url获取数据
// 2. 解析数据
// 3. 统计数据
/*
urlFmt 和 threadNum。
urlFmt 定义了微博实时热度数据的请求 URL，
threadNum 定义了并发请求的线程数。
*/

func main() {

	resultChan := make(chan []byte)
	urlsChan := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < threadNum; i++ {
		wg.Add(1)
		go func() {
			for {
				url, ok := <-urlsChan
				if !ok {
					wg.Done()
					return
				}
				res, err := http.Get(url)
				if err != nil {
					fmt.Println("Error getting data:", err)
				}
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println("Error parsing data:", err)
				}
				resultChan <- data
			}
		}()
	}

	for i := 0; i < 100; i++ {
		url := fmt.Sprintf(urlFmt, i)
		urlsChan <- url
	}

	close(urlsChan)
	wg.Wait()
	close(resultChan)

	results := make(map[string]int)
	for data := range resultChan {
		str := string(data)
		str = str[strings.Index(str, "(")+1 : len(str)-2]
		lines := strings.Split(str, "\n")
		for _, line := range lines {
			if strings.Contains(line, "query") {
				query := strings.Split(strings.Split(line, "'")[1], "'")[0]
				value, err := strconv.Atoi(strings.Split(strings.Split(line, ":")[1], ",")[0])
				if err != nil {
					fmt.Println("Error parsing value:", err)
				}
				results[query] = value
			}
		}
	}

	for query, value := range results {
		fmt.Println(query, value)
	}
}

/*
(1)首先，定义了两个通道，resultChan 和 urlsChan，前者用于存储获取到的数据，后者用于存储请求 URL。然后定义了一个 sync.WaitGroup 变量 wg，用于并发控制。
(2)接下来是一个 for 循环，用于创建 threadNum 个并发请求，每个请求都通过一个匿名函数实现，该匿名函数作为一个 goroutine 运行。在该匿名函数内部，使用无限循环从 urlsChan 中读取 URL，如果读取失败，则调用wg.Done()，表示该线程已经完成了任务并退出。如果 URL 读取成功，向 URL 发送 HTTP GET 请求，获取数据，并将结果存储进 resultChan 中。
(3)接下来是另一个 for 循环，用于向 urlsChan 中发送要爬取的请求 URL，并进行数据请求。创建 threadNum 个 goroutine 的时候往通道里推 URL，来保证请求可以被并发执行。
(4)close(urlsChan) 声明关闭通道以及向所有 wg.Add() 语句记录的线程发送“任务已完成”的信号。wg.Wait() 告诉主线程等待所有 goroutine 完成，并且阻塞程序。
(5)接下来定义了一个空的 map[string]int 类型的变量 results，用于存储结果。下面的代码使用 range 循环读取 resultChan 中的每个元素（即爬取到的数据），并处理数据存储在 results 中。string(data) 用于将 []byte 类型的数据转换为 string；strings.Index() 是一个字符串函数，用于查找某个字符串中特定字符的位置，本例中用于定位括号，以便去除 JSONP 的外围函数；strings.Split() 用于将字符串分割成数组，以便对每一行进行处理；strings.Contains() 用于检查字符串中是否包含指定的子字符串；strings.Split(line, "'")[1] 用于从字符串中提取出第二个单引号之间的内容，本例中用于获取查询词；strings.Split(strings.Split(line, ":")[1], ",")[0] 用于从字符串中提取出逗号之前的数字值，本例中用于获取查询词的热度值。
(6)最后一个 for 循环用于打印结果，将 results 中的每一条数据逐行打印出来。
*/
