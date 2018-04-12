/*
Slice
	1,slice本身是没有数据，是对底层array的一个view
	2,slice可以向后扩展，不可以向前扩展
Reslice
	1,切片重组
*/
package main

import "fmt"

func updateSlice(s []int)  {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1=", s1)
	s2 := arr[:]
	fmt.Println("s2=", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr =", arr)
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Sprintf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s2))
	fmt.Println(s1[3])
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	// s4 and s5 no longer view arr
	fmt.Println("arr =", arr)
}
