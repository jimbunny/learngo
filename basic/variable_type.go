/*
内建变量类型
	bool string
	(u)int, (u)int8, (u)int16, (u)int32, (u),int64, uintptr(指针)
	byte, rune(字符型32）
	float32, float)64, complex64, complex128
强制类型转换
	类型转换是强制的
	var a, b  int = 3, 4
	var c int = math.Sqrt(a*a + b*b)不对
	var c = int(math.Sqrt(float64(a*a + b*b)))对
*/
package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

//欧拉公式 复数
func euler()  {
	c :=3 + 4i
	fmt.Println(cmplx.Abs(c))
	//python cmath: cmath.exp(1j * cmath.pi)+1
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%3f\n",cmplx.Exp(1i * math.Pi) + 1)
}
//勾股定理
func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
//主函数
func main() {
	euler()
	triangle()
}
