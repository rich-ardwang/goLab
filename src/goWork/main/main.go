package main

import (
	"goWork/work"
	"log"
	"sync"
	"time"
)

//names提供了一组用来显示的名字
var names = []string {
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

//namePrinter使用特定方式打印名字
type namePrinter struct {
	name string
}

//Task实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	//time.Sleep(time.Second)
}

var (
	//maxGoRoutines = runtime.NumCPU()
	maxGoRoutines = 16
	maxCounts = 400000
)

//main是所有Go程序的入口
func main() {
	start := time.Now()
	//使用两个goroutine来创建工作池
	p := work.New(maxGoRoutines)

	log.Println(maxGoRoutines)
	var wg sync.WaitGroup
	wg.Add(maxCounts * len(names))

	for i := 0; i < maxCounts; i++ {
		//迭代names切片
		for _, name := range names {
			//创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name : name,
			}

			go func() {
				//将任务提交执行，当Run返回时，
				//我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	//让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()

	cost := time.Since(start)
	log.Println("cost=", cost)
}
