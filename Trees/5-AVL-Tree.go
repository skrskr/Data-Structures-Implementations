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
	Height int `default:"1"`
 }


func (node *TreeNode) Insert(val int) *TreeNode{
	if node == nil {
		return &TreeNode{Val: val, Height: 1}
	}
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
	node.Height = 1 + int(math.Max(float64(node.Left.GetHeight()), float64(node.Right.GetHeight())))

	BalanceFactor := node.GetBalance()

	if BalanceFactor > 1 {
		if val < node.Left.Val {
			return node.RightRotate((node))
		} else {
			node.Left = node.LeftRotate(node.Left)
			return node.RightRotate(node)
		}
	}

	if BalanceFactor < -1 {
		if val > node.Right.Val {
			return node.LeftRotate((node))
		} else {
			node.Right = node.Right.RightRotate(node.Right)
			return node.LeftRotate(node)
		}
	}
	return node
}

func (node *TreeNode) Delete(val int) *TreeNode{
	if node == nil {
		return node
	}
	if (val < node.Val) {
		node.Left = node.Left.Delete(val)
	} else if (val > node.Val) {
		node.Right = node.Right.Delete(val)
	} else {
		temp := nil
		if node.Left == nil {
			temp = node.Right
			node = nil
			return temp
		}
		if node.Right == nil {
			temp = node.Left
			node = nil
			return temp
		}

		temp = node.GetMinValueNode()
		node.key = temp.key
		node.Right = node.Right.Delete(val)
		if node == nil {
			return node
		}

	}
	node.Height = 1 + int(math.Max(float64(node.Left.GetHeight()), float64(node.Right.GetHeight())))

	BalanceFactor := node.GetBalance()

	if BalanceFactor > 1 {
		if val < node.Left.Val {
			return node.RightRotate((node))
		} else {
			node.Left = node.LeftRotate(node.Left)
			return node.RightRotate(node)
		}
	}

	if BalanceFactor < -1 {
		if val > node.Right.Val {
			return node.LeftRotate((node))
		} else {
			node.Right = node.Right.RightRotate(node.Right)
			return node.LeftRotate(node)
		}
	}
	return node
}

// get hieght of node 
func (node *TreeNode) GetHeight() int{ 
	if node == nil {
		return 0
	}
	return node.Height
}

// Get balance of node
func (node *TreeNode) GetBalance() int{ 
	if node == nil {
		return 0
	}
	return node.Left.GetHeight() - node.Right.GetHeight()
}

// Get balance of node
func (node *TreeNode) GetMinValueNode() *TreeNode{ 
	if node == nil || node.Left == nil{
		return node
	}
	return node.Left.GetMinValueNode()
}


// Left rotate
func (node *TreeNode) LeftRotate(node2 *TreeNode) *TreeNode { 
	y := node2.Right
	tmp := y.Left
	y.Left = node2
	node2.Right = tmp

	node2.Height = 1 + int(math.Max(float64(node2.Left.GetHeight()), float64(node2.Right.GetHeight())))
	
	y.Height = 1 + int(math.Max(float64(y.Left.GetHeight()), float64(y.Right.GetHeight())))
	return y
}

// Right rotate
func (node *TreeNode) RightRotate(node2 *TreeNode) *TreeNode { 
	y := node2.Left
	tmp := y.Right
	y.Right = node2
	node2.Left = tmp

	node2.Height = 1 + int(math.Max(float64(node2.Left.GetHeight()), float64(node2.Right.GetHeight())))
	
	y.Height = 1 + int(math.Max(float64(y.Left.GetHeight()), float64(y.Right.GetHeight())))
	return y
}

/* Tree */
type AVLTree struct {
	Root *TreeNode
}


/**
Insert in binary tree

**/

func (tree *AVLTree) Insert(data int) {
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
func (tree *AVLTree) Search(q int) bool {
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

func (tree *AVLTree) Delete(val int) {
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