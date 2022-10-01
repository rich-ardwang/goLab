package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tm := time.Now()
	rand.Seed(tm.Unix()) // Try changing this number!
	answers := []string{
		"Richard",
		"Madaoqi",
		"Pangchengyao",
		"Yeyi",
		"Liulipeng",
		"Wangchao",
	}
	lenth := len(answers)
	for i := 0; i < lenth; i++ {
		ran := rand.Intn(len(answers))
		fmt.Println("Result:", answers[ran])
		answers = append(answers[:ran], answers[ran+1:]...)
	}
}
