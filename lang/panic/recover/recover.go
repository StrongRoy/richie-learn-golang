package main

import (
	"fmt"
)

func tryRecover()  {
	defer func() {
		r := recover()
		if err , ok := r.(error); ok{
			fmt.Println("error occurred:",err)
		}else {
			panic(fmt.Sprintf("I don't know whate to do: %v",r))
		}
	}()
	//panic(errors.New("this is a error"))
	//b := 0
	//a := 5/b
	//fmt.Println(a)
	panic(123)
}


func main() {

	tryRecover()
}
