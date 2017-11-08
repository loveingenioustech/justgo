package main

import (
	"reflect"
	"fmt"
)

type Baby struct {
	name string
	age int
}

func main() {
	var i int
	i = 10
	t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	fmt.Println(t)
	fmt.Println(v)

	// 这段跑不通报panic: reflect: Elem of invalid type
	//b := Baby{"robin", 30}
	//t1 := reflect.TypeOf(b)
	//v1 := reflect.ValueOf(b)
	//
	//tag := t1.Elem().Field(0).Tag  //获取定义在struct里面的标签
	//name := v1.Elem().Field(0).String()  //获取存储在第一个字段里面的值
	//fmt.Println(tag)
	//fmt.Println(name)

	var x float64 = 3.4
	v2 := reflect.ValueOf(x)
	fmt.Println("type:", v2.Type())
	fmt.Println("kind is float64:", v2.Kind() == reflect.Float64)
	fmt.Println("value:", v2.Float())

	// 反射的话，那么反射的字段必须是可修改的
	// 下面这段会报错 panic: reflect: reflect.Value.SetFloat using unaddressable value
	//var x3 float64 = 3.4
	//v3 := reflect.ValueOf(x3)
	//v3.SetFloat(7.1)

	var x4 float64 = 3.4
	p := reflect.ValueOf(&x4)
	fmt.Println(p)
	v4 := p.Elem()
	v4.SetFloat(7.1)
	fmt.Println(x4)
}
