package main

import (
	"fmt"
	"math/rand"
)



func main() {
	demoIf()

	demoGoto()

	demoFor1()

	demoFor2()

	demoBreak()

	demoRange()

	demoSwitch()
}
func demoIf() {
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

func demoGoto() {
	i := 0
Here:   //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	if i<10 {
		goto Here   //跳转到Here去
	}else {
		fmt.Println("finished")
	}

}

func demoFor1(){
	sum := 0;
	for index:=0; index < 10 ; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
}

func demoFor2(){
	// 省略后就类似于while了
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func demoBreak(){
	for index := 10; index>0; index-- {
		if index == 5{
			break // 或者continue
		}
		fmt.Println(index)
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1
}

func demoRange()  {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	for k,v:=range numbers {
		fmt.Println("map's key:",k)
		fmt.Println("map's val:",v)
	}
}


func demoSwitch(){
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	// 可以很多值聚合在了一个case里面，同时，Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	// 但是可以使用fallthrough强制执行后面的case代码。
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

