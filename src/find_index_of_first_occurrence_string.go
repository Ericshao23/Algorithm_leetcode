// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
package main

import "fmt"

// 暴力枚举，依次以原串的每个字符为匹配起始点，匹配需要匹配的字符，一致返回起始点位置，否则继续下一个位置
func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	for i := 0; i <= n-m; i++ {
		j := i
		k := 0
		for k < m {
			if haystack[j] == needle[k] {
				k++
				j++
			}
		}
		if k == m {
			return i
		}
	}
	return -1
}

func strStr_Optimization(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	s1 := "sadbutsad"
	s2 := "sad"
	s3 := "leetcode"
	s4 := "leeto"
	fmt.Println(strStr_Optimization(s1, s2))
	fmt.Println(strStr_Optimization(s3, s4))
}
