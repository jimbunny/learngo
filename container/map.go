/*
Map操作
	1,创建make(map[string]int)
	2,获取元素:m[key]
	3,key不存在时，获得Value类型的初始值
	4,用value, ok := m[key]来判断是否存在key
	5,用delete删除一个key
Map遍历
	1,使用range遍历key，或者遍历key,value对
	2,不好整遍历顺序，如需顺序，需手动对key排序
	3,使用len获得元素个数
Map key
	1,map使用哈希表，必须可以比较相等
	2,除了slice,map,function的内建类型都可以作为key
	3,struct类型不包含上述字段，可以作为key
例题
	寻找最长不含有重复字符的子串 https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
	1,lastOccurred[x]不存在,或者< start -> 无需操作
	2,lastOccurred[x] >= start -> 更新start
	3,更新lastOccurred[x],更新maxLength
*/
package main

import "fmt"

//nonrepeting
func lengthOFNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s){
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
//主函数
func main() {
	m := map[string]string {
		"name": "jim",
		"course": "golang",
		"site": "shanghai",
		"quality": "not bad",
	}
	m2 := make(map[string]int)  // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m{
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if caurseName, ok := m["caurse"]; ok{
		fmt.Println(caurseName)
	}else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println(lengthOFNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOFNonRepeatingSubStr("bbbbbbbb"))
	fmt.Println(lengthOFNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOFNonRepeatingSubStr(""))
	fmt.Println(lengthOFNonRepeatingSubStr("b"))
	fmt.Println(lengthOFNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOFNonRepeatingSubStr("哈哈，你好"))
	fmt.Println(lengthOFNonRepeatingSubStr("一二三二一"))
	fmt.Println(lengthOFNonRepeatingSubStr("黑化肥挥发会发挥会花飞灰化肥挥发发黑会飞花"))
}
