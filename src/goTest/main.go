package main

/*
import (
	"fmt"
	"math/rand"
	"time"
)

//GenRandom 获取随机数
func GenRandom(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end-start) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
			fmt.Println(num)
		}
	}
	return nums
}

func main() {
	nSt := 13500000000
	nEnd := 13999999999
	nCount := 10000000
	//var randInt []int
	GenRandom(nSt, nEnd, nCount)
}
*/
