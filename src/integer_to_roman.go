// 题目链接：https://leetcode.cn/problems/integer-to-roman/
// 输入范围：1～3999
package main

import "fmt"

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "X"},
	{4, "XI"},
	{1, "I"},
}

// 模拟
func intToRoman(num int) string {
	roman := []byte{}
	num = 99
	for _, vs := range valueSymbols {
		fmt.Println(vs)
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	fmt.Println(string(roman))
	return string(roman)
}

// 暴力
var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman_optimization(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

func main() {
	intToRoman(99)
}
