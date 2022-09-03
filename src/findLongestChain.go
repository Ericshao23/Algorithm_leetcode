//https://leetcode.cn/problems/maximum-length-of-pair-chain/

package src

import (
	"math"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	// Slice sorts the slice x given the provided less function.
	// It panics if x is not a slice.
	//
	// The sort is not guaranteed to be stable: equal elements
	// may be reversed from their original order.
	// For a stable sort, use SliceStable.
	//
	// The less function must satisfy the same requirements as
	// the Interface type's Less method.
	//func Slice(x any, less func(i, j int) bool) {
	//	rv := reflectValueOf(x)
	//	swap := reflectSwapper(x)
	//	length := rv.Len()
	//	quickSort_func(lessSwap{less, swap}, 0, length, maxDepth(length))
	//}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	length := len(pairs)
	dp := make([]int, length)
	for i, p := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] {
			if p[0] > q[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[length-1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func findLongestChain_Optimization(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	cur := math.MinInt32
	ans := 0
	for _, p := range pairs {
		if cur < p[0] {
			cur = p[1]
			ans++
		}
	}
	return ans
}

//func main() {
//	pairs := [][]int{
//		{1, 2},
//		{7, 8},
//		{4, 5},
//	}
//	fmt.Println(findLongestChain(pairs))
//}
