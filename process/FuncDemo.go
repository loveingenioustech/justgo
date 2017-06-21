package main

import "fmt"

// 返回a、b中最大值.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//返回 A+B 和 A*B 多个返回值
//func SumAndProduct(A, B int) (int, int) {
//	return A+B, A*B
//}

func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A+B
	Multiplied = A*B
	return
}

func main() {
	demoMax()

	demoMultiReturn()

	demoVaryParas(9,1200,888)
	demoVaryParas(222,666)

	demoPassValue()

	demoPassAddress()

	demoDefer()

	demoType()
}

func demoMax() {
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	//调用函数max(x, y)
	max_xz := max(x, z)
	//调用函数max(x, z)
	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))
	// 也可在这直接调用它
}

func demoMultiReturn()  {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}

func demoVaryParas(arg ...int) {
	// arg ...int告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}


//简单的一个函数，实现了参数+1的操作
func add1(a int) int {
	a = a+1 // 我们改变了a的值
	return a //返回一个新值
}

// 当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，
// 当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。
func demoPassValue() {
	x := 3
	fmt.Println("x = ", x)  // 应该输出 "x = 3"

	x1 := add1(x)  //调用add1(x)

	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出"x = 3"
}

//简单的一个函数，实现了参数+1的操作
func add2(a *int) int { // 请注意，
	*a = *a+1 // 修改了a的值
	return *a // 返回新值
}

// 变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存。只有add2函数知道x变量所在的地址，才能修改x变量的值。
// 所以我们需要将x所在地址&x传入函数，并将函数的参数的类型由int改为*int，即改为指针类型，才能在函数中修改x变量的值。
// 此时参数仍然是按copy传递的，只是copy的是一个指针。
func demoPassAddress(){
	x := 3

	fmt.Println("x = ", x)  // 应该输出 "x = 3"

	x1 := add2(&x)  // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"
	// 传指针的好处
	// 传指针使得多个函数能操作同一个对象。
	// 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
	// Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
}

func demoDefer(){
	// 当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，
	// 在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。
	// 如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

// 在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
func demoType(){
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)

	even := filter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}


