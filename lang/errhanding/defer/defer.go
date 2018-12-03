package main

import (
	"bufio"
	"fmt"
	"os"
	"richie.com/richie/learn_golang/lang/functional/fib"
)

func tryDefer() {
	//defer fmt.Println(1) // defer 相当于又一个栈  先进后出
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occurred")
	//fmt.Println(4)

	for i:=0;i<100;i++{
		defer fmt.Println(i)
		if i == 30{
			panic("printed too many")
		}
	}
}

func writerFile(filename string) {
	//if file, err := os.Create(filename); err != nil {
	if file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666); err != nil {
		if pathError, ok := err.(*os.PathError); !ok{
			panic(err)
		}else{
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,pathError.Path,pathError.Err)

		}

		return
	} else {

		defer file.Close()

		writer := bufio.NewWriter(file)

		defer writer.Flush()

		f := fib.Fibonacci()

		for i := 0; i < 20; i++ {
			fmt.Fprintln(writer, f())
		}

	}
}

func main() {
	writerFile("fib.txt")
}
