package worker

import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	wg sync.WaitGroup
	//工作池
	taskPool chan Worker
}

func New(maxGoroutineNum int) *Pool {
	//1.初始化一个Pool
	p := Pool{
		taskPool: make(chan Worker),
	}

	p.wg.Add(maxGoroutineNum)
	//2.创建maxGoroutineNum个goroutine，并发的丛taskPool中获取任务
	for i := 0; i < maxGoroutineNum; i++ {
		go func() {
			for task := range p.taskPool { //阻塞获取，一旦没有任务，阻塞在这里---无缓冲channel
				//3.执行任务
				task.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交任务到worker池中
func (p *Pool) Run(worker Worker) {
	p.taskPool <- worker
}

func (p *Pool) Shutdown() {
	//关闭通道
	close(p.taskPool)
	p.wg.Wait()
}
