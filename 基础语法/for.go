/*
for
	1,for条件里不需要括号
	2,for的条件里可以省略初始条件，结束条件，递增表达式
	3,省略初始条件，相当于while
	4,什么都省略相当于死循环，没有while
*/
package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)
//将十进制转换成二进制(省略起始条件)
func convertToBin(n int) string {
	result := ""
	for ; n >0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
//读文件（省略起始条件和递增条件）
func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//等同于while循环
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func forever()  {
	for{
		fmt.Println("abc")
	}
}
func main() {
	sum := 0
	for i :=1; i <=100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1011 --> 1101
		convertToBin(234242131),
		convertToBin(0),
		)
	//printFile("abc.txt")
	forever()
}
