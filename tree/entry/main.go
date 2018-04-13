/*
封装
	1,名字一般使用CamelCase
	2,首字母大写:public
	3,首字母小写:private
包
	1,每个目录一个包
	2,main包包含可执行入口
	3,为结构定义的方法必须放在同一个包内
	4,可以是不同的文件
如何扩充系统类型或者别人的类型
	1,定义别名
	2,使用组合
*/
package main

import "fmt"

//结构体方法
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 定义结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

type myTreeNode struct {
	node *treeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.left}
	left.postOrder()
	right := myTreeNode{myNode.node.right}
	right.postOrder()
	myNode.node.print()
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
