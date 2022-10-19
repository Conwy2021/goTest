package main

import (
	"fmt"
	"testing"
)

type sumable1 interface {
	sum(int, int) int
}

type myFunc2 func(int) int

func (f myFunc2) sum(a, b int) int {
	res := a + b
	return f(res)
}

func sum101(num int) int {
	return num * 10
}

func sum1001(num int) int {
	return num * 100
}

func handlerSum1(handler sumable1, a, b int) int {
	res := handler.sum(a, b) //先调用函数的方法，再调用函数自己本身
	fmt.Println(res)
	return res
}

// icansum 实现了 sumable 接口
type icansum struct {
	name string
	res  int
}

func (ics *icansum) sum(a, b int) int {
	ics.res = a + b
	return ics.res
}

func Test45(t *testing.T) {

	newFunc11 := myFunc2(sum101) //可以理解为类型转换
	newFunc21 := myFunc2(sum1001)

	handlerSum1(newFunc11, 1, 1) // 20
	handlerSum1(newFunc21, 1, 1) // 200

	ics := &icansum{"I can sum", 0}
	// 由于 icansum 实现了接口 sumable，所以 handlerSum 可以直接传入 ics 结构体
	handlerSum1(ics, 1, 1) // 2
}
