/*
常量的定义
	1,const filename = "abc.txt"
	2,const数值可作为各种类型使用
	3,const a,b = 3,4
	4,var c = int(math.Sqrt(a*a + b*b))
枚举类型
	1,普通枚举类型
	2,变量自增枚举类型
变量定义要点
	1,变量类型写在变量名之后
	2,编译器可突刺变量类型
	3,没有char,只有rune
	4,原生支持复数类型
*/
package main

import (
	"math"
	"fmt"
)
//常量定义
func consts()  {
	const filename string= "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
//枚举类型
func enums()  {
	const(
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
	const(
		a = iota
		_
		c
		d
		e
	)
	fmt.Println(a, c, d, e)
//	b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	consts()
	enums()
}
