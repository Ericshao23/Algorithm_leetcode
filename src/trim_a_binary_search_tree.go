// https://leetcode.cn/problems/trim-a-binary-search-tree/
package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	for node := root; node.Left != nil; {
		if node.Left.Val < low {
			node.Left = node.Left.Right
		} else {
			node = node.Left
		}
	}
	for node := root; node.Right != nil; {
		if node.Left.Val >= low {
			node.Right = node.Right.Left
		} else {
			node = node.Right
		}
	}
	return root
}
