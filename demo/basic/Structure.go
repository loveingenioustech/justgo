package main

//var和const变量和常量申明
//package和import
//func 用于定义函数和方法
//return 用于从函数返回
//defer 用于类似析构函数
//go 用于并发
//select 用于选择不同类型的通讯
//interface 用于定义接口
//struct 用于定义抽象数据类型
//break、case、continue、for、fallthrough、else、if、switch、goto、default流程控制
//chan用于channel通讯
//type用于声明自定义类型
//map用于声明map类型数据
//range用于读取slice、map、channel数据

// 导入其他的包
import "fmt"

// 常量的定义
const PI = 3.14

// 全局变量
var golbal = "test"

// 一般类型声明
type newType int

// 结构的声明
type s1 struct {

}

// 接口的声明
type golang interface {

}

// 声明一个新的类型
type person struct {
	name string
	age int
}

type Human struct {
	name string
	age int
	weight int
}

type Skills []string

type Student struct {
	Human  // 匿名字段，那么默认Student就包含了Human的所有字段
	Skills // 匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	speciality string
}

type Employee struct {
	Human  // 匿名字段Human
	speciality string
	name string  // 雇员的name字段
}


// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age>p2.age {  // 比较p1和p2这两个人的年龄
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}


// main函数入口
func main() {
	demo1()

	demo2()

	demo3()

	demo4()
}

func demo1(){
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{age:25, name:"Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}

func demo2()  {
	// 我们初始化一个学生
	mark := Student{Human:Human{"Mark", 25, 120}, speciality:"Computer Science"}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)

	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)

	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}

func demo3() {
	// 初始化学生Jane
	jane := Student{Human:Human{"Jane", 35, 100}, speciality:"Biology"}

	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)

	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}


func demo4() {
	Bob := Employee{Human{"Bob in Human", 34, 120}, "Designer", "Bob in Employee"}
	fmt.Println("Bob's name is:", Bob.name)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's human name is:", Bob.Human.name)
}