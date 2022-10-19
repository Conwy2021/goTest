package main

import (
	"fmt"
	"testing"
)

type (
	RegisterService func(native *NativeContract)
	MethodHandler   func(contract *NativeContract) ([]byte, error)
)
type NativeContract struct {
	ContractRef int
	StateDb     string
}

type myFunc1 func(int) int

func _sum10(num int) int {
	fmt.Println(num * 10)
	return num * 10
}

func TestHello2(t *testing.T) {

	newFunc := myFunc1(_sum10)
	newFunc(1)

}

type myInt int

func (mi myInt) IsZero() bool {
	return mi == 0
}

func Test38(t *testing.T) {

	var a myInt
	a = 0
	fmt.Println(a.IsZero())

}

func TestMap(t *testing.T) {
	mapHaiCoder := map[string]string{
		"Server":     "Golang",
		"JavaScript": "Vue",
		"Db":         "Redis",
	}
	value, isOk := mapHaiCoder["Server"]
	fmt.Println("Value =", value, "IsOk =", isOk)
}
