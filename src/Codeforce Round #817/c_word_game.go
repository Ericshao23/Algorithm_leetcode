// // https://codeforces.com/contest/1722/problem/C
// // 使用map，每一个新增的字符都将写入map中，然后遍历每一个人计算其对应分数

// package main

// import "fmt"

// func main() {
// 	var t int
// 	fmt.Scanf("%d\n", &t)
// 	for i := 1; i <= t; i++ {
// 		var n int
// 		mp := make(map[string]int)
// 		//var mp map[string]int
// 		fmt.Scanf("%d\n", &n)
// 		//var a [3][n]string
// 		a := make([][]string, 3)
// 		for j := 0; j < 3; j++ {
// 			a[j] = make([]string, n)
// 			for k := 0; k < n; k++ {
// 				//var str string
// 				//fmt.Scan(&a[j][k])
// 				fmt.Scanf("%s", &a[j][k])
// 				//a[i][j] = str
// 				//fmt.Println("_", len(a[i][j]), "_", a[i][j], "_")
// 				mp[a[j][k]]++
// 			}
// 		}
// 		for j := 0; j < 3; j++ {
// 			total := 0
// 			for k := 0; k < n; k++ {
// 				if mp[a[j][k]] == 1 {
// 					total += 3
// 				} else if mp[a[j][k]] == 2 {
// 					total++
// 				}
// 			}
// 			fmt.Printf("%d ", total)
// 		}
// 		fmt.Printf("\n")
// 	}
// }