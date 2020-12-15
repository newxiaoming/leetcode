package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res = []int{}
	fmt.Println(root)
	if root != nil {
		res = append(res, root.Val)
		for _, node := range root.Children {
			res = append(res, preorder(node)...)
		}
	} else {
		return []int{}
	}

	return res

}

func main() {

	node2 := &Node{
		Val:      2,
		Children: nil,
	}
	node3 := &Node{
		Val:      3,
		Children: nil,
	}
	node4 := &Node{
		Val:      4,
		Children: nil,
	}
	node5 := &Node{
		Val:      5,
		Children: nil,
	}
	node6 := &Node{
		Val:      6,
		Children: nil,
	}

	root := &Node{
		Val:      1,
		Children: nil,
	}
	root.Children = append(root.Children, node3)
	node3.Children = append(node3.Children, node5)
	node3.Children = append(node3.Children, node6)
	root.Children = append(root.Children, node2)
	root.Children = append(root.Children, node4)
	rst := preorder(root)
	fmt.Println(rst)
}
