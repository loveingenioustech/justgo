package main

import "fmt"

func main() {
	// map也就是Python中字典的概念，它的格式为map[keyType]valueType
	// map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。

	// var numbers map[string]int
	// 另一种map的声明方式, 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据

	// map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	// 初始化一个字典
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")  // 删除key为C的元素
	fmt.Println(rating)

	// map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了
	fmt.Println(m["Hello"])
}
