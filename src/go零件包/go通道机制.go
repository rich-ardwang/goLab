package main

import (
	"fmt"
	"time"
)
func delay(i int) {
	for a:=i; a>0; a-- {
		if a==1 {
			fmt.Printf("s=%d sec.\n", a)
		} else {
			fmt.Printf("s=%d sec...\n",a)
		}
		time.Sleep(1*time.Second)
	}
}

func main() {
	//通道缓冲元素个数是2，说明当向通道中放入超过2个int时，通道阻塞
	ch := make(chan int, 2)
	fmt.Println("main1")
	go func() {
		fmt.Println("go1")
		ch<- 99
		ch<- 18
		ch<- 1
		//通道阻塞，当99被读取时，才能放入3
		ch<- 3
		//通道阻塞，当18被读取时，才能继续执行
		fmt.Println("go2")
	}()
	fmt.Println("main2")
	delay(5)
	r := <-ch
	fmt.Printf("r=%d\n",r)
	delay(5)
	h := <-ch
	fmt.Printf("h=%d\n",h)
	delay(60*60)
}


