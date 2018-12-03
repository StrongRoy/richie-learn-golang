# if

if contents, err := ioutil.ReadFile(filename); err != nil {
    fmt.Println("cannot print file contents",err)
}else {
	fmt.Printf(string(contents))
}

* if条件可以赋值
* if的条件里赋值的变量作用域就在这个if语句里


# switch
 
func grade(score int) string {
  	g := ""
  	switch {
  	case score < 0 || score > 100:
  		panic(fmt.Sprintf("Wrong score: %d", score))
  	case score < 60:
  		g = "F"
  	case score < 80:
  		g = "C"
  	case score < 90:
  		g = "B"
  	case score < 100:
  		g = "A"
  	}
  	return g
  }
  
 # for 
 func convertToBin(n int) string {
 	result := ""
 	for ; n > 0; n /= 2 {
 		lsb := n % 2
 		result = strconv.Itoa(lsb) + result
 	}
 	return result
 }
 
 func printFile(filename string)  {
 	file ,err := os.Open(filename)
 	if err != nil{
 		panic(err)
 	}
 	scanner := bufio.NewScanner(file)
 	for scanner.Scan(){
 		fmt.Println(scanner.Text())
 	}
 }
 
 func forever()  {
 	for {
 		fmt.Println("abc")
 	}
 }
 
# 函数
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

* 函数返回多个值时候可以起名字
* 仅用于非常简单的函数
* 对于调用这而言没有区别


语法要点
* 返回值类型写在最后
* 可返回多个值
* 函数作为参数
* 没有默认参数，可选参数

# 参数传递  值传递? 引用传递
* go语言只有**值传递**一种方式

