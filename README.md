# Algorithm_leetcode

## [29. Divide Two Integers](https://leetcode.cn/problems/divide-two-integers/)

### 思路：

1. 单纯的暴力利用减法代替除法，循环利用被除数减去除数，直到被除数小于除数，当然这是两个数同为正的情况下。这里需要对负数进行特判（**同号为正，异号为负，对最高位进行异或来判断结果是否为负数**），其次对于两个数异号，我们需要将其转换为同号，但是考虑到为32位整数，转为正数会溢出，所以异号情况统一转换为负数。
2. 优化思路，在单纯的减法上，试图减去除数的倍数。
3. 继续优化，利用位运算，提高运算速度。
### [code](./src/divide_two_numbers.go)

## [12. Integer to Roman](https://leetcode.cn/problems/integer-to-roman/)

### 思路
1. 模拟，根据罗马数字的唯一表示法，为了表示一个给定的整数 `num`，我们寻找不超过 `num` 的最大符号值，将 `num` 减去该符号值，然后继续寻找不超过 `num` 的最大符号值，将该符号拼接在上一个找到的符号之后，循环直至 `num` 为 0。
2. 暴力，找规律，整理每个位置对应数字的字符。
### [code](./src/integer_to_roman.go)

## [13.Roman to Integer](https://leetcode.cn/problems/roman-to-integer/)

## 思路
1. 模拟：一般情况下罗马数字中的小数字在大数字的右边，若数字的右边比左边的数字的大则该数字取反。
2. 暴力：从右往左便利字符串，遇到比当前字符串大的就加上其对应值，遇到比当前字符串小的就减去其对应值。
### [code](./src/roman_to_integer.go)


## [1475. Final Prices With a Special Discount in a Shop](https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/)

### 思路
1. 暴力：从第 i+1 件商品开始依次向后遍历，直到找到第一个满足 `prices[j] <= prices[i]`的下标 j 即可求出该物品的最终折扣价格。
2. 单调栈——使得查询`prices` 中每个元素对应位置的右边的第一个更小的元素值时不需要再遍历 `prices`。
### [code](./src/finalPrices.go)


## [646.Maximum Length of Pair Chain](https://leetcode.cn/problems/maximum-length-of-pair-chain/)

### 思路
1. 动态规划：定义`dp[i]`为以`pair[i]`为结尾的最长数对链长度。计算`dp[i]`时，可以先找出所有满足`pair[i][0]>pair[j][1]`的j，并求出最大的`dp[j]`,`dp[i]`的值就可以被赋予这个最大值+1，要求计算`dp[i]` 时，所有潜在的`dp[j]` 已经计算完成，可以先将 `pairs` 进行排序来满足这一要求。初始化时，`dp` 需要全部赋值为 1。
2. 贪心：为了使得数对链路最长，需要在选择是一个数对时，最优的就是挑选第二个数字最小的；同样选择第二个数对时也是在剩下数对中选择第二个数字最小的且第一个数字大于上一个数对第二个数字的数对。
### [code](./src/findLongestChain.go)


## [1582. Special Positions in a Binary Matrix](https://leetcode.cn/problems/special-positions-in-a-binary-matrix/)
### 思路
1. 模拟：枚举每一个位置，按照特殊位置的定义去判断是否为特殊位置（即对应行和对应列除该元素外均为0）；**但是矩阵中的每一个元素只能为1/0，所以可以通过计算每一行和列的元素和来得到每一行中的一的个数**
2. 列的标记值：定义第 j 列的标记值为：该列所有 1所在行中的 1的数量之和，用原始矩阵的第一行来作为我们标记列的额外空间，从而降低空间复杂度

### [code](./src/special_positions_in_a_binary_matrix.go)

## [43. Multiply Strings](https://leetcode.cn/problems/multiply-strings/)
### 思路
1. 竖式计算：将`num2`的每一位和`num1`相乘，结果累加
### [code](./src/multiply_strings.go)

## [652. Find Duplicate Subtrees](https://leetcode.cn/problems/find-duplicate-subtrees/)
### 思路
1.将每一棵子树都「序列化」成一个字符串，并且保证：
+ 相同的子树会被序列化成相同的子串；
+ 不同的子树会被序列化成不同的子串。
2. 使用三元组进行唯一表示：用一个三元组直接表示一棵子树，即 (x, l, r)它们分别表示：
+ 根节点的值为 x； 
+ 左子树的序号为 l； 
+ 右子树的序号为 r

### [code](./src/find_duplicate_subtress.go)

## [JZ37 序列化二叉树](https://www.nowcoder.com/practice/cf7e25aa97c04cc1a68c8f040e71fb84?tpId=265&tqId=39239&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D13&difficulty=undefined&judgeStatus=undefined&tags=&title=)
+ 二叉树序列化(Serialize)：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。
+ 二叉树反序列化(Deserialize):根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。
### 思路
1. DFS（深度优先搜索）：
   1. [先序遍历](./src/serialize_deserialize.go)
   2. [中序遍历](./src/serialize&deserialize_optimization.go)

## [828.Count Unique Characters of All Substrings of a Given String](https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/)
### 思路
+ 当字符是"头元素"，那么出现次数为`数组长度 - 元素下标位置`
+ 当字符是"尾元素"，那么出现次数为`元素下标位置 - （-1）`
+ 字符是"中间元素"，出现次数为：`(元素下标位置 - (-1)) * (数组长度 - 元素下标位置)`

## [1592. Rearrange Spaces Between Words](https://leetcode.cn/problems/rearrange-spaces-between-words/)
### 思路
1. 首先按照空格对字符串分割，得到单词集合和空格数
   1. 单词书为1，则将全部的空格拼接到这个单词后面即可
   2. 计算出单词间的间隔，并按照单词及间隔来进行拼接，若拼接后仍有多余的空格，则将剩下的空格拼接在末尾即可。
### [code]()