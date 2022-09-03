package main

func numSpecial(mat [][]int) int {
	rowsum := make([]int, len(mat))
	colsum := make([]int, len(mat[0]))
	for i, row := range mat {
		for j, x := range row {
			rowsum[i] += x
			colsum[j] += x
		}
	}
	ans := 0
	for i, row := range mat {
		for j, x := range row {
			if x == 1 && rowsum[i] == 1 && colsum[j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func numSpecial_optimization(mat [][]int) (ans int) {
	for i, row := range mat {
		temp := 0
		for _, x := range row {
			temp += x
		}
		if i == 0 {
			temp--
		}
		if temp > 0 {
			for j, x := range row {
				if x == 1 {
					mat[0][j] += temp
				}
			}
		}
	}
	for _, x := range mat[0] {
		if x == 1 {
			ans++
		}
	}
	return ans
}
