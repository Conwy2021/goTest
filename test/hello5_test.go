package main

import (
	"fmt"
	"math/big"
	"testing"
)

func Test1034(t *testing.T) {
	//左看右闭合
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[1:9:10]
	fmt.Println(s1)
	s2 := s[1:]
	fmt.Println(s2)
	s3 := s[1:4]
	fmt.Println(s3)
	s4 := s[:2]
	fmt.Println(s4)

}

func Test202211181804(t *testing.T) {
	a := big.NewInt(1)
	b := big.NewInt(2)

	c1 := a.Cmp(b)
	fmt.Println(c1)

}
