package main

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	for i, p := range prices {
		dicount := 0
		for _, q := range prices[i+1:] {
			if q <= p {
				dicount = q
				break
			}
		}
		ans[i] = p - dicount
	}
	return ans
}

func finalPrices_Optimization(prices []int) []int {
	n := len(prices)
	ans := make([]int, n)
	st := []int{0}
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(st) > 1 && st[len(st)-1] > p {
			st = st[:len(st)-1]
		}
		ans[i] = p - st[len(st)-1]
		st = append(st, p)
	}
	return ans
}
