package main

import (
	"fmt"
	"math"
)

/* Tree Node */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode

}


func (node *TreeNode) isLeaf() bool {
	return node.Left == nil && node.Right == nil
}


func (node *TreeNode) Insert(val int) {
	if node == nil {
		return
	} else if (val <= node.Val) {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val, Left: nil, Right: nil}
		} else {
			node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val, Left: nil, Right: nil}
		} else {
			node.Right.Insert(val)
		}
	}
}



/* Tree */
type Tree struct {
	Root *TreeNode

}

func (tree *Tree) IsEmpty() bool {
	return tree.Root == nil
}

// Tree Length
func (tree *Tree) Length() int {
	return calculateTreeLength(tree.Root) + 1
}

func calculateTreeLength(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return 1 + calculateTreeLength(root.Right)
	}

	if root.Right == nil {
		return 1 + calculateTreeLength(root.Left)
	}

	return 2 + int(math.Max(float64(calculateTreeLength(root.Left)), float64(calculateTreeLength(root.Right))))
}

// inset in tree

func (tree *Tree) Insert(data int) *Tree {
	if tree.Root == nil {
		tree.Root = &TreeNode{Val: data, Left: nil, Right: nil}
	} else {
		tree.Root.Insert(data)
	}
	return tree
}


func main() {
	// tree := &Tree{}
	// root := &TreeNode{Val: 3}
	// tree.Root = root
	// node1 := &TreeNode{Val: 1}
	// node2 := &TreeNode{Val: 4}
	// node3 := &TreeNode{Val: 5}
	// tree.Root.Left = node1
	// tree.Root.Right = node2
	// tree.Root.Left.Left = node3

	// fmt.Println(tree.Root.Left.isLeaf())
	// fmt.Println(tree.Length())

	tree := &Tree{}
	tree.Insert(100).
		Insert(-20).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(15).
		Insert(5).
		Insert(-10)

	fmt.Println(tree.IsEmpty())
	fmt.Println(tree.Length())
}