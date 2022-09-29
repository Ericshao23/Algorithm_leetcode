package main

import (
	"fmt"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
next:
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if s1[(i+j)%len(s1)] != s2[j] {
				continue next
			}
		}
		return true
	}
	return false
}

func isFlipedString_optimization(s1 string, s2 string) bool {
	// 将字符s1拼接欸s1+s1，那么无论如何旋转，s2一定是s1+s1的子字符串
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Println(isFlipedString(s1, s2))
	fmt.Println(isFlipedString_optimization(s1, s2))
}
