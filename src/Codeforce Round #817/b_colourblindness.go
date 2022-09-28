// // https://codeforces.com/contest/1722/problem/B
// // 将第一行字符中的所有'B'全部替换为字符'C'，然后比较两个字符是否一致

// package main

// import "fmt"

// func main() {
// 	var t int
// 	fmt.Scanf("%d\n", &t)
// 	for i := 0; i < t; i++ {
// 		var n int
// 		//fmt.Scanf("%d\n", &n)
// 		fmt.Scan(&n)
// 		var str1, str2 string
// 		//fmt.Scanf("%s\n", &str1)
// 		fmt.Scan(&str1)
// 		b1 := []byte(str1)
// 		//fmt.Scanf("%s\n", &str2)
// 		fmt.Scan(&str2)
// 		b2 := []byte(str2)
// 		for j := 0; j < n; j++ {
// 			if b1[j] == 'B' {
// 				b1[j] = 'G'
// 			}
// 			if b2[j] == 'B' {
// 				b2[j] = 'G'
// 			}
// 		}
// 		str1 = string(b1)
// 		str2 = string(b2)
// 		//fmt.Println(str)
// 		//fmt.Println(str2)
// 		if str2 == str1 {
// 			fmt.Println("YES")
// 		} else {
// 			fmt.Println("NO")
// 		}
// 	}
// }