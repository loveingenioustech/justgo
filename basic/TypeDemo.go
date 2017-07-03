package main

import "fmt"

// 设置别名
type (
	byte int8
	rune int32
	文本 string
)

func main() {
	var b 文本
	b = "中文类型名"
	fmt.Println(b)
}
