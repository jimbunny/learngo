/*
函数
	1,返回多个值时可以起名字，仅用于非常简单的函数，对于调用者而言没有区别
	2,函数式编程，函数式一等公民,函数可以作为参数
	3,可变参数作为参数列表（没有默认参数，可选参数）
	4,返回值类型写在最后面
*/
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//switch
func eval1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return  a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}
//双返回值 带余除法 13 / 3 = 4 ... 1
func div(a, b int) (q, r int) {
	//q = a / b
	//r = a % b
	return a / b, a % b
}
//函数式编程
func apply(op func(int, int) int, a, b int)  int {
	//reflect反射获取op的值
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}
//可变参数作为参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}
//主函数
func main() {
	//fmt.Println(eval1(4,4,"/"))
	if result, err := eval1(3,4,"x");err != nil {
		fmt.Println("Error:", err)
	}else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
