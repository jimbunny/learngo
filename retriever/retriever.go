/*
接口的定义
	1,由使用者定义
接口的实现
	1,接口的实现是隐式的
	2,只要实现接口里定义的方法
接口变量
	1,接口的变量自带指针
	2,接口变量同样采用值传递，几乎不需要使用接口的指针
	3,指针接受者实现只能以指针的方式使用;值接受者都可
	4,表示任何类型：interface{}
	5,Type Assertion
	6,Type Switch
*/
package main

import (
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"fmt"
	"time"
	"os"
	"io"
	"bufio"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func downland(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster)  {
	poster.Post("http://www.imooc.com",map[string]string{
		"name": "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url  = "http://www.imooc.com"
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake www.baidu.com"}
	r = &retriever
	fmt.Printf("%T %v\n",r, r)
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut: time.Minute,
	}

	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok{
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a seesion")
	fmt.Println(session( &retriever ))
	printFile("abc.txt")
	s := `abc"d"
	kkkkk
	123

	p`
	fmt.Println(s)
	//fmt.Println(downland(r))
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func inspect(r Retriever){
	fmt.Printf("%T %v\n",r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Printf("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
