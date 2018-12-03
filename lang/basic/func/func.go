package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupperted operation: %s ", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d,%d)", opName, a, b)
	return op(a, b)
}

func sum(number ...int) int {
	s := 0
	for i := range number{
		s += number[i]
	}
	return s
}
func pow(a, b int)  int {
	return int(math.Pow(float64(a),float64(b)))
}


func swap(a,b *int)  {
	*b, *a = *a,*b
}

func main() {

	if restlt, err := eval(3, 4, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("error: %s ", restlt)
	}

	//fmt.Println(eval(2, 4, "x"))

	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a),float64(b)))
	},3,4))

	fmt.Println(sum(1,2,3,4,5,6,7,8,9))

	a,b := 3,4
	swap(&a,&b)
	fmt.Println(a,b)
}
