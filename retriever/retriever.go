/*
接口的定义
	1,由使用者定义
接口的实现
	1,接口的实现是隐式的
	2，只要实现接口里定义的方法
*/
package main

import (
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
} 
func downland(r Retriever) string {
	return r.Get("http://www.baidu.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake www.baidu.com"}
	r = real.Retriever{}
	fmt.Println(downland(r))
}
