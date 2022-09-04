// https://codeforces.com/contest/1722/problem/A
// 只需要判断升序排列后是否与Timur升序排列后是否一致

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Timru是对字符Timur的升序排列
	var name = "Timru"
	// 对name字符串排序
	// 类似c++的sort，升序排列
	//b := []byte(name)
	//sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	//name = string(b)
	//fmt.Println(name)
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scanf("%d\n", &m)
		var s string
		fmt.Scanf("%s\n", &s)
		b1 := []byte(s)
		sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
		s = string(b1)
		//fmt.Println(s)
		//fmt.Println(i)
		if name == s {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
