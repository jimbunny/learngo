/*
使用 var 关键字
	1,var a, b, c bool
	2,s1, s2 string = "hello:, :world"
	3,可放在函数内，或直接放在包内
	4,使用var()集中定义变量
使用:=定义变量
	1,a, b, s1, s2 := true, false, 3, "hello", "world"
	2,只能在函数内使用
*/
package main

import "fmt"

//包内部变量，不可以使用:=
var aa = 3
var ss = "kkk"
var bb = true

//或者
var (
	cc = 3
	dd = "kkk"
	ff = true
)

//定义变量
func variableZeroValue() {
	var a int
	var b string
	fmt.Println(a, b)
	fmt.Printf("%d %s\n", a, b)
}

//定义变量付初值
func variableInitialValue() {
	var a, s int = 3, 4
	var b string = "abc"
	fmt.Println(a, b, s)
}

//自动推断变量类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//简写声明定义变量
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//主函数
func main() {
	fmt.Println("hello word")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
}
