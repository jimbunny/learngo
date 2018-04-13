/*
面向对象
	1,go语言仅支持封装，不支持继承和多态
	2,go语言没有class,只有struct
结构的创建
	1,不论地址还是结构体本身，一律使用.来访问成员
	2,使用自定义工厂函数
	3,注意返回了局部变量的地址
结构定义方法
	1,显示定义和命名方法接受者
使用指针作为方法的接受者
	1,只有使用指针才可以改变结构体内容
	2,nil指针也可以调用方法
值接受者vs指针接受者
	1,要改变内鹅绒必须使用指针接受者
	2,结构过大也考虑使用指针接受者
	3,一致性：如有指针接受者，最好都是指针接受者
	4,值接受者是go语言特有
	5,值/指针接受者均可接收值/指针
*/
package main

import (
	"fmt"
)

// 定义结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

//结构体方法
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

//主函数
func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.traverse()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()
	fmt.Println()

	root.right.left.setValue(4)
	root.right.left.print()

	root.print()
	root.setValue(100)

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()

	var tRoot *treeNode
	tRoot.setValue(200)
	tRoot = &root
	tRoot.setValue(300)
	tRoot.print()
}
