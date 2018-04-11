/**
if
	1,if的条件里不需要括号
	2,if的条件里可以赋值
	3,if的条件里赋值的变量作用域就在这个if语句里
switch
	1,switch会自动break,除非使用fallthrough
	2,switch后面可以没有表达式
 */
package main

import (
	"fmt"
	"io/ioutil"
)
//if条件无括号
func bounded(v int) int {
	if v > 100{
		return 100
	}else if v < 0{
		return 0
	}else {
		return v
	}
}
//switch 无表达式
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}
func grade(score int) string {
	g := ""
	switch {
	case score > 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
func main() {
	fmt.Println(bounded(101))
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("%s\n", contents)
	}
	//if的条件里可以赋值
	if contents, err := ioutil.ReadFile(filename);err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("%s\n", contents)
	}
	fmt.Println(eval(1, 2 , "+"))
	fmt.Println(grade(0),grade(50),grade(60),grade(80),grade(100),grade(101),)
}
