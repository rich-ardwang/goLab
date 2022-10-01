package main

import (
	"flag"
	"fmt"
)

var mRet = make(map[int][][]int)
var mPos = make(map[int]int)
var cnt int = 0
func main() {
	var nQueenCnt int
	flag.IntVar(&nQueenCnt, "qc", 8, "Set queen count")
	flag.Parse()
	GotoNextLine(1, nQueenCnt)
	fmt.Printf("【%d】皇后问题：总共找到【%d】种解决方案。\n", nQueenCnt, cnt)
	fmt.Println("下面打印出第一种和最后一种：")
	OutputGrh(mRet[1])
	fmt.Println("--------------------------华丽分割线--------------------------")
	OutputGrh(mRet[cnt])
}
func GotoNextLine(line int, nQueenCnt int) bool {
	col := 1
	for ; col <= nQueenCnt; col++ {
		mPos[line] = col
		if IsAttack(line, col) {
			continue
		} else {
			if nQueenCnt == line {
				cnt++
				d := make([][]int, nQueenCnt)
				copy(d, mPos)
				mRet[cnt] = d
				continue
			} else {
				if (!GotoNextLine((line+1), nQueenCnt)) {
					continue
				} else {
					return true
				}
			}
		}
	}
	if 1 == line && col > nQueenCnt {
		return true
	}
	return false
}
func IsAttack(line int, col int) bool {
	if line == 1 {
		return false
	}
	lineB := line - 1
	height := 1
	for i := lineB; i >= 1; i-- {
		colB := mPos[i]
		if colB == col || colB == (col + height) || colB == (col - height) {
			return true
		}
		height++
	}
	return false
}
func copy(des [][]int, src map[int]int) {
	for i:=0; i<len(src); i++ {
		des[i] = append(des[i], i+1)
		des[i] = append(des[i], src[i+1])
	}
}
func OutputResult(dt [][]int) {
	for _, v := range dt {
		fmt.Printf("line = %d, col = %d\n", v[0], v[1])
	}
}
func OutputGrh(dt [][]int) {
	l1 := "\t "
	l2 := "\t "
	for i:=0; i<len(dt); i++ {
		l1 += fmt.Sprintf("%d ", i+1)
		l2 += "_ "
	}
	fmt.Println(l1)
	fmt.Println(l2)
	for i:=0; i<len(dt); i++ {
		l := fmt.Sprintf("%d\t|", i+1)
		for j:=0; j<len(dt); j++ {
			if (j+1) == dt[i][1] {
				l += "*|"
			} else {
				l += "_|"
			}
		}
		fmt.Println(l)
	}
}