# 面向对象
    * go语言仅支持封装，不支持继承和多态
    * go语言没有class，只有struct



# 结构的创建
    var root treeNode
	root = treeNode{value:3}
	fmt.Println(root)
	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new(treeNode)

* 无论地址还是结构本身 ，一律使用.来访问成员
```go
func createNode(value int) *treeNode  {
	return &treeNode{value:value}  // 在c++中这李会返回的是局部变量的地址
}

root.right.left = createNode(2)
```

* 使用自定义工厂函数
* 注意返回了局部变量的地址！ 


# 结构创建在堆上还是栈上？
    在其他语言中如c++需要devaloper将局部变量分配到栈上，函数一旦退出，自动销毁，
    如果要传出需要在堆上创建，创建完成之后需要开发者手动销毁
    在java中几乎所有的东西都分配在堆上 使用new
    
**go语言不需要知道创建在堆上还是栈上，由编译器和运行环境决定**


# 只接受着vs指针接受者
* 要改变内容必须使用指针接受者
* 结构过大也考虑使用指针接受者
* 一致性： 如果由指针接受者最好用指针接受者
    
* **值接受者**是go语言特有的
* 值/指针接受者均可接受值/指针
    

# 封装
* 名字一般使用CamelCase
* 首字母大写：public    --》 针对的是包
* 首字母小写：private

# 包

* 每个目录一个包
* main包包含可执行入口
* 为结构定义的方法必须全部放在一个包内
* 可以是不同文件

go语言的所有的参数都是传值



# 扩展已有的类型
# 如果扩充系统类型加或别人的类型
* 定义别名
* 使用组合

