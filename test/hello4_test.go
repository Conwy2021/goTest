package main

import (
	"fmt"
	"testing"
)

type a interface {
	sum1(int, int) int
	sum2(int) int
}

//以下是定义一个函数类型b
type b func(int) int

//针对这个函数类型可以再定义方法，如：
func (f b) b15(a, b int) int {
	res := a + b
	return f(res)
}
func (f b) b14(a int) int {
	return a
}

//形参指定传入参数为函数类型b
func Oper(fub b, a int) int {
	return fub(a)
}

//使用定义的函数类型
func Test1440(t *testing.T) {
	oper := Oper(sumtest, 2)
	fmt.Println(oper)

}

func sumtest(num int) int {

	return num + 2
}

func sum1011(num int) int { //自定义函数
	return num * 10
}

func sum10012(num int) int { //自定义函数
	return num * 100
}

func handlerSum11(handler sumable1, a, b int) int { //自定义函数
	res := handler.sum(a, b)
	fmt.Println(res)
	return res
}

// c 实现了 a 接口
type c struct {
	name string
	res  int
}

func (ics *c) sum1(a, b int) int { //c 的方法 实现了type a 接口
	ics.res = a + b
	return ics.res
}

func (ics *c) sum2(a1 int) int { //c 的方法 实现了type a 接口

	return a1
}

func test11() int { //函数返回函数

	return sum1011(1)
}

func test12(_c c) int { //函数返回方法

	return _c.sum2(_c.res)
}

func Test1431(t *testing.T) {
	_c := c{}
	i := test12(_c)
	fmt.Println(i)

}

func Test51(t *testing.T) {
	c2 := &c{"a", 1} //表示这种数字类型的变量
	c3 := c{"b", 2}  //表示一种数字类型
	fmt.Println(c2)
	fmt.Println(c3)
	sum := c2.sum2(1)
	fmt.Println(sum)
	sum1 := c3.sum1(1, 3)
	fmt.Println(sum1)
}

func Test46(t *testing.T) {

	//c:=&c{"hello",1}
	e := &c{"hello3", 1}
	fmt.Println(e)
	c1 := c{}
	c1.name = "hello"
	c1.res = 1
	fmt.Println(c1)
	d := new(c)
	d.name = "hello2"
	fmt.Println(d)

	fmt.Println(111)

}
