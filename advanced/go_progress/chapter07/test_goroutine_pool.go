package main

import (
	"fmt"
	"time"
)

type Task struct {
	F func()
}

// NewTask 创建任务
func NewTask(f func()) *Task {
	task := Task{F: f}
	return &task
}

// TaskRun 任务执行
func (t *Task) TaskRun() {
	t.F()
}

// GoroutinePool 协程池
type GoroutinePool struct {
	CapNum int
	// 进任务的管道
	InChannel chan *Task
	// 任务调度的管道
	WorkChannel chan *Task
}

func NewPool(cap_num int) *GoroutinePool {

	pool := GoroutinePool{
		CapNum:      cap_num,
		InChannel:   make(chan *Task),
		WorkChannel: make(chan *Task),
	}

	return &pool
}

// TaskInChannelOut 从InChannel管道拿到任务，放到WorkChannel
func (p *GoroutinePool) TaskInChannelOut() {
	for task := range p.InChannel {
		p.WorkChannel <- task
	}
}

// Worker 任务执行者从WorkChannel获取任务并执行
func (p *GoroutinePool) Worker() {

	// WorkChannel

	for task := range p.WorkChannel {
		task.TaskRun()
		fmt.Println("任务执行完毕")
	}
}

func (p *GoroutinePool) PoolRun() {

	// WorkChannel获取任务，并任务执行
	for i := 0; i < p.CapNum; i++ {
		go p.Worker() // 开启指定数量的协程执行任务
	}

	// 从InChannel管道拿到任务，本质是往WorkChannel里面添加任务
	p.TaskInChannelOut()

	close(p.WorkChannel)
	close(p.InChannel)

}

// 1:从InChannel获取任务并写入WorkChannel

// 2.从WorkChannel里面获取任务并执行

// 少了往InChannel写入任务

func main() {

	cap_num := 5
	//task := NewTask(func() {
	//	fmt.Println(time.Now())
	//})

	pool := NewPool(cap_num)

	go func() {

		for {
			task := NewTask(func() {
				fmt.Println(time.Now())
			})
			pool.InChannel <- task
		}
	}()

	// 任务调度
	pool.PoolRun()

}
