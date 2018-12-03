package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value} // 在c++中这李会返回的是局部变量的地址
}

// 接受者
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = Value
}

