package main

import "fmt"

func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		x3, x5, x7 := dp[p3]*3, dp[p5]*5, dp[p7]*7
		dp[i] = min(x3, x5, x7)
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
		if dp[i] == x7 {
			p7++
		}
	}
	return dp[k]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}

	} else if b < c {
		return b
	} else {
		return c
	}
}

func main() {
	fmt.Println(getKthMagicNumber(5))
}
