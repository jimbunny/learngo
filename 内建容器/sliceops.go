/*
Slice
	1,s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
	2,添加元素时如果超越cap,系统会重新分配更大的底层数组
	3,由于值传递的关系，必须接受append的返回值 s = append(s, val)
*/
package main

import "fmt"

//打印结构
func printSlice(s []int)  {
	fmt.Printf("%v,len=%d, cap=%d\n",s, len(s), cap(s))
}
func main() {
	fmt.Println("Creating slice")
	var s []int // zero value for slice is nil
	for i :=0; i < 100; i++{
		printSlice(s)
		s = append(s, 2 * i+ 1)
	}
	fmt.Println(s)

	s6 := []int{2, 4, 6, 8}
	printSlice(s6)

	s7 := make([]int, 16)

	s8 :=make([]int, 10, 32)
	printSlice(s7)
	printSlice(s8)
	// 拷贝
	fmt.Println("Copying slice")
	copy(s7, s6)
	printSlice(s7)
	// 删除中间元素
	fmt.Println("Deleting elements from slice")
	s7 = append(s7[:3], s7[4:]...)
	printSlice(s7)
	// 删除首个元素
	fmt.Println("Popping from front")
	front := s7[0]
	s7 = s7[1:]
	fmt.Println(front)
	// 删除末尾元素
	fmt.Println("Popping from back")
	tail := s7[len(s7) -1]
	s7 = s7[:len(s7) - 1]
	fmt.Println(tail)
	printSlice(s7)
}
