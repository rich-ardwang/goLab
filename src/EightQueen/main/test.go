package main

import "fmt"

func main() {
	s1 := make([]int, 3, 5)
	s2 := make([]int, 2)
	fmt.Println(s1)
	fmt.Println(s2)
	s1 = append(s1, 5)
	s1 = append(s1, 6)
	s1 = append(s1, 7)
	s2 = append(s2, 5)
	s2 = append(s2, 6)
	s2 = append(s2, 7)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}

