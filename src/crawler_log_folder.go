// https://leetcode.cn/problems/crawler-log-folder/

package main

func minOperations(logs []string) (depth int) {
	for _, log := range logs {
		if log == "./" {
			continue
		}
		if log != "../" {
			depth++
		} else if depth > 0 {
			depth--
		}
	}
	return
}
