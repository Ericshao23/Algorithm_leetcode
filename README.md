# Algorithm_leetcode

## [29. Divide Two Integers](https://leetcode.cn/problems/divide-two-integers/)

### 思路：

1. 单纯的暴力利用减法代替除法，循环利用被除数减去除数，直到被除数小于除数，当然这是两个数同为正的情况下。这里需要对负数进行特判（**同号为正，异号为负，对最高位进行异或来判断结果是否为负数**），其次对于两个数异号，我们需要将其转换为同号，但是考虑到为32位整数，转为正数会溢出，所以异号情况统一转换为负数。
2. 优化思路，在单纯的减法上，试图减去除数的倍数
3. 继续优化，利用位运算，提高运算速度


## [12. Divide Two Integers](https://leetcode.cn/problems/integer-to-roman/)

### 思路
1. 模拟，根据罗马数字的唯一表示法，为了表示一个给定的整数 `num`，我们寻找不超过 `num` 的最大符号值，将 `num` 减去该符号值，然后继续寻找不超过 `num` 的最大符号值，将该符号拼接在上一个找到的符号之后，循环直至 `num` 为 0。
2. 暴力，找规律，整理每个位置对应数字的字符

## [1475. Final Prices With a Special Discount in a Shop](https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/)
### 思路
1. 暴力：从第 i+1 件商品开始依次向后遍历，直到找到第一个满足 `prices[j] <= prices[i]`的下标 j 即可求出该物品的最终折扣价格
2. 单调栈——使得查询`prices` 中每个元素对应位置的右边的第一个更小的元素值时不需要再遍历 `prices`



