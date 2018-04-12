/*
指针
	1,指针不能运算
参数传递
	值传递？引用传递？
	1,go语言只有值传递一种方式（python，java内置类型值传递，自定义类型引用传递（for example:int））
*/
package main

import "fmt"
//值传递
func swap(a, b int)  {
	b, a = a, b
}
//引用传递
func swap1(a, b *int)  {
	*b, *a = *a, *b
}
func main() {
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)
	swap1(&a, &b)
	fmt.Println(a, b)
}
