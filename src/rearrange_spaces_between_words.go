// https://leetcode.cn/problems/rearrange-spaces-between-words/
package main

import "strings"

func reorderSpaces(text string) (ans string) {
	words := strings.Fields(text)
	space := strings.Count(text, " ")
	l := len(words) - 1
	if l == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/l)) + strings.Repeat(" ", space%l)
}
