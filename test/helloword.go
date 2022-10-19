package main

import (
	"fmt"
	"github.com/jacksonyoudi/gomodone"
	"goTest25/test2" //引入的是文件
	"rsc.io/quote"
)

func main() {
	fmt.Println("go mod test start...")
	defer fmt.Println("go mod test end!")
	fmt.Println(quote.Hello())
	fmt.Println(gomodone.SayHi("Roberto"))
	Hello2()
	test3.Test2() //引用其他包的方法 大写是public 方法 使用的是包名
	//test2.test2() 小写就报错  是私有方法
	stooges := []string{"Moe", "Larry", "Curly"}
	lang := []string{"php", "golang", "java", "aaa"}
	stooges = append(stooges, lang...)
	fmt.Println(stooges) //[Moe Larry Curly php golang java]
}
