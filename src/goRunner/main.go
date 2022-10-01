package main

import (
	"goRunner/runner"
	"log"
	"os"
	"time"
)

//timeout规定了必须在多少秒内处理完成
const timeout = 5 * time.Second

//main是程序的入口
func main() {
	log.Println("Starting work.")

	//为本次执行分配超时时间
	r := runner.New(timeout)

	//加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	//执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

//createTask返回一个根据id休眠指定秒数的事例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
