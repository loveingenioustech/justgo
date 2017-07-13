package main

import "fmt"

func main() {
	// 长度也是数组类型的一部分，因此[3]int与[10]int是不同的类型，数组也就不能改变长度。
	// 数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。
	// 如果要使用指针，那么就需要用到slice类型了。
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 指定位置
	d := [3]int{2:1}
	fmt.Println(d)

	// 二维数组
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println(doubleArray)
	fmt.Println(easyArray)

	// 动态数组 Slice
	// slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
	slice := []byte {'a', 'b', 'c', 'd'} // 和声明array一样，只是少了长度
	fmt.Println(slice)
}
