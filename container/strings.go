/*
rune
	1,使用range遍历ops,rune对
	2,使用utf8.RuneCountInString获得字符数量
	3,使用len获得字节长度
	4,使用[]byte获得字节
其他字符串操作
	1,Fields,Split,Join
	2,Contains,Index
	3,ToLower,ToUpper
	4.Trim,TrimRight,TrimLeft
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes呵呵哒！" // UTF-8 英文1字节，中文3字节
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
