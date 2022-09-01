// 题目地址：https://leetcode.cn/problems/divide-two-integers/
// tips：被除数和除数均为32位有符号整数，除数不为0
package main

import "math"

// 单纯的暴力
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}
	res := 0
	for dividend <= divisor {
		dividend -= divisor
		res++
	}
	return sign * res
}

// 暴力的优化，减去除数的倍数
func divide_optimization(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}
	res := 0
	for dividend <= divisor {
		value, k := divisor, 1
		//0xc000000是十进制-2^30的十六进制表示
		//判断value >= 0xc0000000, 保证value+value不会溢出
		for value >= 0xc0000000 && dividend <= value+value {
			value += value
			// 如果k已经大于最大值一半，直接返回最小值，因为k+=k会溢出
			if k > math.MinInt32/2 {
				return math.MinInt32
			}
			k += k
		}
		dividend -= value
		res += k
	}
	return sign * res
}

// 利用位运算提高运算速度
func divide_continue_optimization(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	res := 0
	// 处理边界情况
	if divisor == math.MinInt32 {
		if dividend == divisor {
			return 1
		} else {
			return 0
		}
	}

	// 被除数先减去一个除数
	if dividend == math.MinInt32 {
		dividend -= -abs(divisor)
		res += 1
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	dividend = abs(dividend)
	divisor = abs(divisor)
	for i := 31; i >= 0; i-- {
		if (dividend>>i)-divisor >= 0 {
			dividend = dividend - (divisor << i)
			if res > math.MaxInt32-(1<<i) {
				return math.MinInt32
			}
			res += 1 << i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
