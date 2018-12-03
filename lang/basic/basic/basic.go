package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 我们需要使用fmt包中的Println()函数

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func varabelInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func varableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func varableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {

	fmt.Printf("%.3f\n ",
		cmplx.Exp(1i*math.Pi)+1)
	//cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a,b))
}

func calcTriangle(a,b int) int  {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	// b kb mb gb tb pb

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello, world. 你好，世界!")
	variableZeroValue()
	varabelInitialValue()
	varableTypeDeduction()
	varableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	euler()
	enums()
}
