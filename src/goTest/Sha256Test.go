package main

/*
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const CharacterSet string = "0123456789abcdefghijklmnopqrstuvwxyz~!@#$%^&*()_+-={}|[]:;,.<>?"

var cnt int64 = 0

func main() {
	var sha256 []string
	sha256 = calcSha256InputByOutput(4, 1)
	shaLen := len(sha256)
	fmt.Println("Now output result...")
	for i := 0; i < shaLen; i++ {
		fmt.Printf("%s \n", sha256[i])
	}
}

func calcSha256(input string) string {
	var sRet string = ""
	h := sha256.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)
	sRet = hex.EncodeToString(bs)
	return sRet
}

func getChar(index int) string {
	var sRet string = ""
	if (index < 0) || (index > len(CharacterSet)) {
		return sRet
	}
	c := []byte(CharacterSet)
	sRet = string(c[index])
	return sRet
}

func getString(index []int) string {
	var sRet string = ""
	indexLen := len(index)

	//Check parameter
	if indexLen <= 0 {
		return sRet
	} else {
		for i := 0; i < indexLen; i++ {
			if (index[i] < 0) || (index[i] > len(CharacterSet)) {
				return sRet
			}
		}
	}

	for i := 0; i < indexLen; i++ {
		sRet += getChar(index[i])
	}
	return sRet
}

func zeroCmpare(sha256Res string, zeroCnt int) bool {
	sha256Slice := []byte(sha256Res)
	for i := 0; i < zeroCnt; i++ {
		if 0 != sha256Slice[i] {
			return false
		}
	}
	return true
}

func getSha256Output(preIndex []int, calcbit int, zeroCnt int) []string {
	var sRets []string
	if calcbit < 0 {
		return sRets
	}
	charCnt := len(CharacterSet)
	if 0 == calcbit {
		for i := 0; i < charCnt; i++ {
			cnt++
			fmt.Printf("run...%d\n", cnt)
			preIndex = append(preIndex, i)
			tryS := getString(preIndex)
			resS := calcSha256(tryS)
			if zeroCmpare(resS, zeroCnt) {
				ret := "input:[" + tryS + "], Sha256Output:[" + resS + "]."
				sRets = append(sRets, ret)
			}
		}
	} else {
		calcbit--
		for i := 0; i < charCnt; i++ {
			preIndex = append(preIndex, i)
			sRets = getSha256Output(preIndex, calcbit, zeroCnt)
		}
	}
	return sRets
}

func calcSha256InputByOutput(calcCnt int, zeroCnt int) []string {
	fmt.Println("calc start...")
	var sRets []string
	for i := 0; i < calcCnt; i++ {
		fmt.Printf("calc count %d...\n", i+1)
		var preIndex []int
		sRets = getSha256Output(preIndex, i, zeroCnt)
	}
	fmt.Println("calc finish.")
	return sRets
}
*/
