//https://www.nowcoder.com/practice/cf7e25aa97c04cc1a68c8f040e71fb84?tpId=265&tqId=39239&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D13&difficulty=undefined&judgeStatus=undefined&tags=&title=
package main

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
func serialize(root *TreeNode) string {
	ans := "["

	// 中序遍历
	nodeQueue := make([]*TreeNode, 0)
	nodeQueue = append(nodeQueue, root) // 添加根元素

	for len(nodeQueue) != 0 {
		tmpNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if tmpNode == nil {
			ans += "null" + ","
			continue // null无需无差别添加子节点
		} else {
			ans += strconv.Itoa(tmpNode.Val) + ","
		}

		nodeQueue = append(nodeQueue, tmpNode.Left)
		nodeQueue = append(nodeQueue, tmpNode.Right)
	}

	ans = ans[:len(ans)-1] // 去掉最后一个","

	return ans + "]"
}

func deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1]          // 去除[]
	dataSplit := strings.Split(data, ",") // 切割
	if len(dataSplit) == 1 {              // 按照序列化规则，如果根节点有内容，至少切出三个元素
		return nil
	}
	nodeQueue := make([]*TreeNode, 0)
	dataSplitCount := 0
	rootVal, _ := strconv.Atoi(dataSplit[dataSplitCount]) // Itoa 不会出错，Atoi可能会出错，所以有error
	root := &TreeNode{Val: rootVal, Left: nil, Right: nil}
	dataSplitCount++
	nodeQueue = append(nodeQueue, root) // 添加首元素

	for len(nodeQueue) != 0 {
		tmpNode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		leftNodeContext := dataSplit[dataSplitCount] // 添加左节点
		dataSplitCount++
		if leftNodeContext == "null" {
			tmpNode.Left = nil
		} else {
			leftVal, _ := strconv.Atoi(leftNodeContext)
			leftNode := &TreeNode{Val: leftVal, Left: nil, Right: nil}
			tmpNode.Left = leftNode
			nodeQueue = append(nodeQueue, leftNode)
		}

		rightNodeContext := dataSplit[dataSplitCount]
		dataSplitCount++
		if rightNodeContext == "null" {
			tmpNode.Right = nil
		} else {
			rightVal, _ := strconv.Atoi(rightNodeContext)
			rightNode := &TreeNode{Val: rightVal, Left: nil, Right: nil}
			tmpNode.Right = rightNode
			nodeQueue = append(nodeQueue, rightNode)
		}
	}
	return root
}
