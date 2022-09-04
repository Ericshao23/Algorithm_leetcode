// https://codeforces.com/contest/1717/problem/A
package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int32
		fmt.Scan(&n)
		fmt.Println(n + 2*(n/2+n/3))
	}
}
