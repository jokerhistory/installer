package main

//https://studygolang.com/articles/464
import (
	"os"
	"text/template"
)

//http://blog.csdn.net/sryan/article/details/52353937
func main() {
	name := "waynehu"
	tmpl, err := template.New("test").Parse("hello,111 {{2.}}222") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, name) //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
