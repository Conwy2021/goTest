package main

import (
	"fmt"
	"testing"
)

type myFunc3 func(int) int

func (f myFunc3) sum3(a, b int) int {
	res := a + b
	return f(res)
}

func sum10111(num int) int {
	return num * 10
}

func handlerSum2(handler myFunc3, a, b int) int {
	res := handler.sum3(a, b) //先调用函数的方法，再调用函数自己本身
	fmt.Println(res)
	return res
}

func Test1644(t *testing.T) {
	newFunc11 := myFunc3(sum10111) //可以理解为类型转换
	handlerSum2(newFunc11, 1, 1)   // 20
}

//使用外部的函数做为自己的函数的一部分 算是一种解耦吧  并且可以有效的管理函数内的参数作用范围
