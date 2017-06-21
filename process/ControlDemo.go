package main

import (
	"fmt"
	"math/rand"
)



func main() {
	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := computedValue(); x > 10 {
		fmt.Println(x, "x is greater than 10")
	} else {
		fmt.Println(x, "x is less than 10")
	}

	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	// fmt.Println(x)

}

func computedValue() int {
	return rand.Intn(30)
}
