package main

import (
	"fmt"
	"testing"
)

type myFunc func(int) int

func (f myFunc) sum(a, b int) int {
	res := a + b
	return f(res)
}

func sum10(num int) int {
	return num * 10
}

func Sum10(num int) int {
	return num * 10
}

func sum100(num int) int {
	return num * 100
}

func handlerSum(handler myFunc, a, b int) int {
	res := handler.sum(a, b)
	fmt.Println(res)
	return res
}

func Test1503(t *testing.T) {

	newFunc2 := myFunc(sum100)
	fmt.Println(newFunc2)
	fmt.Printf("%T\n", newFunc2)
	fmt.Printf("%T\n", sum100)
}

func Test29(t *testing.T) {
	//newFunc1 := myFunc(sum10)
	newFunc2 := myFunc(sum100)
	fmt.Print(newFunc2)
	handlerSum(sum10, 1, 1)    // 20
	handlerSum(newFunc2, 1, 1) // 200
	a := Sum10(3)
	fmt.Println(a)
}
