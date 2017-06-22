package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width*r.height
}

type Circle struct {
	radius float64
}

// 定义method
// 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
// method里面可以访问接收者的字段
// 调用method通过.访问，就像struct里面访问字段一样
func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}


func main() {
	demoProcedure()

	demoObject1()

	demoObject2()

	demoObject3()

	demoObject4()
}
func demoProcedure() {
	// 面向过程的实现
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))
}

func demoObject1()  {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}

func demoObject2() {
	boxes := BoxList {
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}


// 因为同级目录中已经有Human、Student、Employee，所以这里用小写
type human struct {
	name string
	age int
	phone string
}

type student struct {
	human //匿名字段
	school string
}

type employee struct {
	human //匿名字段
	company string
}

//在human上面定义了一个method
func (h human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//student的method重写human的method
func (e *student) SayHi() {
	fmt.Printf("Hi, I am %s, I study at %s. Call me on %s\n", e.name,
		e.school, e.phone) //Yes you can split into 2 lines here.
}

func demoObject3() {
	mark := student{human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := employee{human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}

func demoObject4() {
	mark := student{human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := employee{human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}