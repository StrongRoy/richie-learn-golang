package main

import "fmt"

var lastOccurred = make([]int,0xffff)

func lengthofNonReeatingSubStr(s string) int {
	//lastOccurred := make(map[rune]int)

	// stores last occurred pos +1.
	//0 means not seen.

	for i := range lastOccurred{
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI:= lastOccurred[ch]; lastI >= start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}
func main() {
	fmt.Println(lengthofNonReeatingSubStr("abcabcbb"))
	fmt.Println(lengthofNonReeatingSubStr("bbbbb"))
	fmt.Println(lengthofNonReeatingSubStr("pwwkew"))
	fmt.Println(lengthofNonReeatingSubStr(""))
	fmt.Println(lengthofNonReeatingSubStr("b"))
	fmt.Println(lengthofNonReeatingSubStr("abcedf"))
	fmt.Println(lengthofNonReeatingSubStr("这里是北京"))
	fmt.Println(lengthofNonReeatingSubStr("一二三二一"))
	fmt.Println(lengthofNonReeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
