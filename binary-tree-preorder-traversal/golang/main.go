package main

import (
	"fmt"
)

// 144.二叉树的前缀遍历
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	fmt.Println(root)
	res = append(res, root.Val)
	if root.Left != nil {
		fmt.Println("root(left) ", root.Val)
		res = append(res, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		fmt.Println("root(right)", root.Val)
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res

}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	rst := preorderTraversal(root)
	fmt.Println(rst)
}
