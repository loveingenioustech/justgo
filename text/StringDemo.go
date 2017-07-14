package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	demoContains()

	demoJoin()

	demoIndex()

	demoRepeats()

	demoReplace()

	demoSplit()

	demoTrim()

	demoFields()

	demoStrconv()
}

func demoContains(){
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

// 字符串链接，把slice a通过sep链接起来
func demoJoin()  {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}

// 在字符串中查找sep所在的位置，返回位置值，找不到返回-1
func demoIndex(){
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}

// 	重复s字符串count次，最后返回重复的字符串
func demoRepeats(){
	fmt.Println("ba" + strings.Repeat("na", 2))
}

// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
func demoReplace(){
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}

// 把s字符串按照sep分割，返回slice
func demoSplit(){
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

// 在s字符串的头部和尾部去除cutset指定的字符串
func demoTrim(){
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
}

// 去除s字符串的空格符，并且按照空格分割返回slice
func demoFields(){
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}

// 字符串转换
func demoStrconv(){
	// 添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	// Parse 系列函数把字符串转换为其他类型
	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1)
}

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}
