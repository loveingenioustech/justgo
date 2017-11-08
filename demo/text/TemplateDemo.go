package main

import (
	"html/template"
	"os"
	"strings"
	"fmt"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	demoTemplate1()

	demoTemplate2()

	demoTemplate3()

	demoTemplate4()

	//demoTemplate5()

	demoTemplate6()
}

func demoTemplate1()  {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Robin"}
	t.Execute(os.Stdout, p)
}

func demoTemplate2(){
	f1 := Friend{Fname: "nan"}
	f2 := Friend{Fname: "yexiaowei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	p := Person{UserName: "robin",
		Emails:  []string{"robin@126.com", "robin@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

}

// 注意：if里面无法使用条件判断，例如.Mail=="robin@gmail.com"，这样的判断是不正确的，if里面只能是bool值
func demoTemplate3()  {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

// 自定义模板函数
func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func demoTemplate4(){
	f1 := Friend{Fname: "nan"}
	f2 := Friend{Fname: "sean"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
				{{range .Emails}}
					an emails {{.|emailDeal}}
				{{end}}
				{{with .Friends}}
				{{range .}}
					my friend name is {{.Fname}}
				{{end}}
				{{end}}
				`)
	p := Person{UserName: "Robin",
		Emails:  []string{"robin@126.com", "robin@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}


// 模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
func demoTemplate5(){
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}

// 嵌套模板
func demoTemplate6(){
	s1, _ := template.ParseFiles("text/header.tmpl", "text/content.tmpl", "text/footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}
