package main

import (
	"fmt"
)

/* Tree Node */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func (node *TreeNode) Insert(val int) {

	if (val <= node.Val) {
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
		} else {
			node.Left.Insert(val)	
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
		} else {
			node.Right.Insert(val)
		}
	}
}


/* Tree */
type Tree struct {
	Root *TreeNode
}


/**
Insert in binary tree

**/

func (tree *Tree) Insert(data int) {
	if tree.Root == nil {
		tree.Root = &TreeNode{Val: data, Left: nil, Right: nil}
	} else {
		tree.Root.Insert(data)
	}
}

/**
Search in binary tree
**/


// Search on tree
func SearchBinaryTree(root *TreeNode, q int) bool {
	if root == nil {
		return false
	}

	if root.Val == q {
		return true
	}

	if q > root.Val {
		return SearchBinaryTree(root.Right, q)
	}
	return SearchBinaryTree(root.Left, q)
}

// Search on tree
func (tree *Tree) Search(q int) bool {
	return SearchBinaryTree(tree.Root, q)
}

func printPreOrder(n *TreeNode) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Val)
		printPreOrder(n.Left)
		printPreOrder(n.Right)
	}
}

/**
Delete from tree
*/

func MinValued(root *TreeNode) *TreeNode {
	temp := root
	for temp != nil && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func (tree *Tree) Delete(val int) {
	DeleteBST(tree.Root, val)
}

func DeleteBST(root *TreeNode, val int) *TreeNode{
	if root == nil {
		return root
	}
	
	if root.Val > val {
		root.Left = DeleteBST(root.Left, val)
	} else if (root.Val < val) {
		root.Right = DeleteBST(root.Right, val)
	} else {
		// No children
		if root.Right == nil && root.Left == nil {
			root = nil 
			return root
		}
		// no left children but rigth
		if root.Left == nil && root.Right != nil {
			temp := root.Right
			root = nil
			root = temp
			return root
		}
		// no Right children but left exists
		if root.Left != nil && root.Right == nil {
			temp := root.Left
			root = nil
			root = temp
			return root
		}
		// node has left and right
		tempNode := MinValued(root.Right)
		root.Val = tempNode.Val
		root.Right = DeleteBST(root.Right, tempNode.Val)
	}

	return root
}


func main() {
	tree := &Tree{}
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(11)
	fmt.Println("Pre-order============")
	printPreOrder(tree.Root)
	fmt.Println("Delete - result============")
	tree.Delete(5)
	fmt.Println("Search - result============")
	printPreOrder(tree.Root)
	// fmt.Println(tree.Search(11))
}