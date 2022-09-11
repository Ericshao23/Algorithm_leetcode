// https://codeforces.com/contest/1728/problem/A
// 最大值cnt的出现位置就是答案
package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		maxn := math.MinInt
		temp := math.MinInt
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
			if a[j] >= maxn {
				maxn = a[j]
				temp = j
			}
		}
		fmt.Println(temp + 1)
	}
}
