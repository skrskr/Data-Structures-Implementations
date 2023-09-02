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

// insert in tree

func (tree *Tree) Insert(data int) *Tree {
	if tree.Root == nil {
		tree.Root = &TreeNode{Val: data, Left: nil, Right: nil}
	} else {
		tree.Root.Insert(data)
	}
	return tree
}

/**
Tree Traversal - Preorder

1- Visit root node
2- Visit all the nodes in the left subtree
3- Visit all the nodes in the right subtree
*/

func PreOrder(node *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return
	}

	fmt.Println(node.Val)
	PreOrder(node.Left)
	PreOrder(node.Right)

}

/**
Tree Traversal - Inorder

1- First, visit all the nodes in the left subtree
2- Then the root node
3- Visit all the nodes in the right subtree
*/


func InOrder(node *TreeNode) {
	if node == nil {
		return
	}

	InOrder(node.Left)
	fmt.Println(node.Val)	
	InOrder(node.Right)
}

/**
Tree Traversal - Postorder

1- Visit all the nodes in the left subtree
2- Visit all the nodes in the right subtree
3- Visit the root node
*/


func PostOrder(node *TreeNode) {
	if node == nil  {
		return
	}
	
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Val)
}


func main() {
	tree := &Tree{}
	tree.Insert(100).
		Insert(-20).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(15).
		Insert(5).
		Insert(-10)

	fmt.Println("PRE - Order")
	PreOrder(tree.Root)
	fmt.Println("IN - ORDER")
	InOrder(tree.Root)
	fmt.Println("POST - ORDER")
	PostOrder(tree.Root)
}