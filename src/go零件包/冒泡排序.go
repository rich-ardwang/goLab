package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetInputD(n int) []int {
	t := time.Now()
	r := rand.New(rand.NewSource(t.Unix()))
	s := r.Perm(n)
	return s
}

func SortD(s []int, b bool) []int {
	l := len(s)

	//数组复制方式一
	//r := make([]int, l, l)   //为r分配len为l,cap为l，和s一样大小
	//copy(r, s)   //实现数组复制，如果是r=s，则r不过是s的引用而已，如果是r=s[:]，则r是s的一个切片，在底层实际上还是对s的部分引用

	//数组复制方式二
	r := append([]int{}, s[:]...)

	if l > 2 {
		var br bool = false
		for br == false {
			br = true
			for i := 0; i < (l - 1); i++ {
				if b == true {
					if r[i] > r[i+1] {
						swap(&r[i], &r[i+1])
						br = false
					}
				} else {
					if r[i] < r[i+1] {
						swap(&r[i], &r[i+1])
						br = false
					}
				}
			}
		}
	}
	return r
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	s := GetInputD(99)
	fmt.Printf("s=%v\n",s)
	r := SortD(s, false)
	f := SortD(s, true)
	fmt.Printf("r=%v\n",r)
	fmt.Printf("f=%v\n",f)
	fmt.Printf("s=%v\n",s)
}

