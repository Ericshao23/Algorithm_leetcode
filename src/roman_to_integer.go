// https://leetcode.cn/problems/roman-to-integer/
package main

import "fmt"

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	n := len(s)
	ans := 0
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

func romanToInt_optimization(s string) int {
	n := len(s)
	ans := 0
	var max int
	for i := n - 1; i >= 0; i-- {
		value := symbolValues[s[i]]
		if max <= value {
			ans += value
			max = value
		} else {
			ans -= value
		}
	}
	return ans
}

func main() {
	fmt.Println(romanToInt_optimization("III"))
}
