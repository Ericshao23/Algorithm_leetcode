package main

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, str *string) {
	if root == nil {
		*str += "null"
		return
	}
	// 先序遍历
	*str += strconv.Itoa(root.Val) + ","
	dfs(root.Left, str)
	dfs(root.Right, str)
	return
}

// Encodes a tree to a single string.
func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := ""
	dfs(root, &ans)
	return ans
}

// Decodes your encoded data to tree.
func Deserialize(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	str := strings.Split(s, ",")
	return buildTree(&str)
}

// build binary tree
func buildTree(str *[]string) *TreeNode {
	if (*str)[0] == "null" {
		*str = (*str)[1:] //没办法go是值传递，你需要改变切片的长度，让构造可以进行下去
		return nil
	}
	nums, _ := strconv.Atoi((*str)[0])
	*str = (*str)[1:]
	root := &TreeNode{Val: nums}
	if len(*str) != 0 {
		root.Left = buildTree(str)
	}
	if len(*str) != 0 {
		root.Right = buildTree(str)
	}
	return root
}

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));
