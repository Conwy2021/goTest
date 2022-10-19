package main

import (
	"fmt"
	"testing"
)

func add(a, b int) int     { return a + b }
func sub(c int, d int) int { return c - d }
func mul(e int, f int) int { return e * f }

func Test1559(t *testing.T) {

	fmt.Printf("%T\n", add) //%T的意思 打印值的类型
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", mul)
	add := add(1, 1)
	sub := sub(1, 1)
	mul := mul(1, 1)
	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)

}
