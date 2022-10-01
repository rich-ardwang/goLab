package main

import (
	"fmt"
	"sync"
	"time"
)

var a int = 0
var m sync.Mutex

func main() {
	go func() {
		for {
			m.Lock()
			a++
			fmt.Println("w1: ", a)
			m.Unlock()
			time.Sleep(5*time.Second)
		}
	}()

	go func() {
		for {
			m.Lock()
			a++
			fmt.Println("w2: ", a)
			m.Unlock()
			time.Sleep(5*time.Second)
		}
	}()

	go func() {
		for {
			m.Lock()
			a++
			fmt.Println("w3: ", a)
			m.Unlock()
			time.Sleep(5*time.Second)
		}
	}()

	go func() {
		for {
			m.Lock()
			a++
			fmt.Println("w4: ", a)
			m.Unlock()
			time.Sleep(5*time.Second)
		}
	}()

	go func() {
		for {
			m.Lock()
			a++
			fmt.Println("w5: ", a)
			m.Unlock()
			time.Sleep(5*time.Second)
		}
	}()

	for {
		fmt.Println("main...")
		time.Sleep(5*time.Second)
	}
}
