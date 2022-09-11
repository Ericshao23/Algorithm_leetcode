// https://codeforces.com/contest/1728/problem/B
package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			a[j] = j + 1
		}
		for t := n & 1; t < n-2; t += 2 {
			temp := a[t]
			a[t] = a[t+1]
			a[t+1] = temp
		}
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", a[j])
		}
		fmt.Printf("\n")
	}
}
