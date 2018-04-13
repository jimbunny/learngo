/*
数组
	1,数量写在类型前
	2,数组的遍历
	3,可通过_省略变量
	4,不仅range,任何地方都可通过_省略变量
	5,如果只要i，可写成for i := range numbers
	6,数组是值类型
	7,[10]int和[20]int是不同类型
	8,调用func f(arr [10]int)会拷贝数组
	9,在go语言中一般不直接使用数组，使用切片
range
	为什么要range
	1,意义明确，美观。

*/
package main

import "fmt"

//值传递
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

//引用传递
func printArray1(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	(*arr)[0] = 100
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	fmt.Println("printArray(arr3)")
	printArray1(&arr3)
	fmt.Println("printArray(arr1)")
	printArray1(&arr1)
	fmt.Println(arr1, arr3)
}
