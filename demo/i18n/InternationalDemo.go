package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

func main() {
	demoText()

	demoDate()

	demoMoney()
}

var locales map[string]map[string]string

func demoText(){
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en

	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn
	
	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))

	en["how old"] ="I am %d years old"
	cn["how old"] ="我今年%d岁了"

	fmt.Printf(msg(lang, "how old"), 30)
	fmt.Println()

	lang = "en"
	fmt.Printf(msg(lang, "how old"), 30)
	fmt.Println()
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}

func demoDate()  {
	en := make(map[string]string, 10)
	cn := make(map[string]string, 10)
	en["time_zone"]="America/Chicago"
	cn["time_zone"]="Asia/Shanghai"

	en["date_format"]="%Y-%m-%d %H:%M:%S"
	cn["date_format"]="%Y年%m月%d日 %H时%M分%S秒"

	locales["en"] = en
	locales["zh-CN"] = cn

	lang := "zh-CN"
	loc,_:=time.LoadLocation(msg(lang,"time_zone"))

	t:=time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(date(msg(lang,"date_format"),t))

	lang = "en"
	fmt.Println(t.Format(time.RFC3339))
}

func date(fomate string,t time.Time) string{
	fmt.Println(fomate)
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	//解析相应的%Y %m %d %H %M %S然后返回信息
	fomate = strings.Replace(fomate, "%Y", strconv.Itoa(year), -1)
	fomate = strings.Replace(fomate, "%m", month.String(), -1)
	fomate = strings.Replace(fomate, "%d", strconv.Itoa(day), -1)
	fomate = strings.Replace(fomate, "%H", strconv.Itoa(hour), -1)
	fomate = strings.Replace(fomate, "%M", strconv.Itoa(min), -1)
	fomate = strings.Replace(fomate, "%S", strconv.Itoa(sec), -1)

	return fomate
}

func demoMoney(){
	en := make(map[string]string, 10)
	cn := make(map[string]string, 10)

	en["money"] ="USD %d"
	cn["money"] ="￥%d元"

	locales["en"] = en
	locales["zh-CN"] = cn

	lang := "zh-CN"
	fmt.Println(money_format(msg(lang,"money"),100))

}

func money_format(fomate string,money int64) string{
	return fmt.Sprintf(fomate,money)
}