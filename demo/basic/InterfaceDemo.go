package main

import (
	"fmt"
	"strconv"
)

type Human1 struct {
	name string
	age int
	phone string
}

type Student1 struct {
	Human1 //匿名字段
	school string
	loan float32
}

type Employee1 struct {
	Human1 //匿名字段
	company string
	money float32
}

//Human1实现SayHi方法
func (h Human1) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human1实现Sing方法
func (h Human1) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee1重载Human1的SayHi方法
func (e Employee1) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human1,Student1和Employee1实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	demoInterface1()

	demoInterface2()

	demoInterface3()

	demoInterface4()

	demoInterface5()
}


func demoInterface1() {
   // 通过上面的代码，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
	// Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。
	mike := Student1{Human1{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student1{Human1{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee1{Human1{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee1{Human1{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student1
	i = mike
	fmt.Println("This is Mike, a Student1:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee1
	i = tom
	fmt.Println("This is tom, an Employee1:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}

func demoInterface2()  {
	// 空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
	// 空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，
	// 因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。

	// 定义a为空接口
	var a interface{}
	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	fmt.Println(a)

	a = s
	fmt.Println(a)
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human1) String() string {
	// 任何实现了String方法的类型都能作为参数被fmt.Println调用
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func demoInterface3()  {
	// interface的变量可以持有任意实现该interface类型的对象
	Bob := Human1{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func demoInterface4()  {
	list := make(List, 3)
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

func demoInterface5(){
	// `element.(type)`语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用`comma-ok`。
	list := make(List, 3)
	list[0] = 1 //an int
	list[1] = "Hello" //a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list{
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}