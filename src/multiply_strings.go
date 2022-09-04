// https://leetcode.cn/problems/multiply-strings/
package main

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	//str1 := []byte(num1)
	//str2 :=[]byte(num2)
	m, n := len(num1), len(num2)
	ans := make([]int, m+n)
	for i := n - 1; i >= 0; i-- {
		x := int(num2[i]) - '0'
		for j := m - 1; j >= 0; j-- {
			y := int(num1[j] - '0')
			ans[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		ans[i-1] += ans[i] / 10
		ans[i] %= 10
	}
	res := ""
	idx := 0
	if ans[0] == 0 {
		idx = 1
	}
	for ; idx < n+m; idx++ {
		//将整数转换为10进制字符串
		// Itoa is equivalent to FormatInt(int64(i), 10).
		// FormatInt returns the string representation of i in the given base,
		// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
		// for digit values >= 10.
		//func FormatInt(i int64, base int) string {
		//	if fastSmalls && 0 <= i && i < nSmalls && base == 10 {
		//	return small(int(i))
		//}
		//	_, s := formatBits(nil, uint64(i), base, i < 0, false)
		//	return s
		//}
		res += strconv.Itoa(ans[idx])
	}
	return res
}
