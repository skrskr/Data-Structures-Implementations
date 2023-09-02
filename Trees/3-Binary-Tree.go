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


// Tree Traversal
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



/**
Full Binary Tree

A full Binary tree is a special type of binary tree in which every parent node/internal node has either two or no children.

It is also known as a proper binary tree.
*/

func IsFullBinaryTree(node *TreeNode) bool {

	if node == nil {
		return true
	}

	if node.Left == nil && node.Right == nil {
		return true
	}

	if node.Left != nil && node.Right != nil {
		return (IsFullBinaryTree(node.Left) && IsFullBinaryTree(node.Right))
	}
	
	// left is nil or right is nil
	return false
}


/**
Prefect Binary Tree

A perfect binary tree is a type of binary tree in which every internal node has exactly two child nodes and all the leaf nodes are at the same level.
*/

func CalculateTreeDepth(node *TreeNode) int {
	depth := 0

	for node != nil {
		node = node.Left
		depth += 1
	}
	return depth
}

func IsPerfectBinaryTree(node *TreeNode,depth int, level int) bool {

	if node == nil {
		return true
	}

	if node.Left == nil && node.Right == nil {
		return ( depth == level + 1)
	}

	if node.Left == nil || node.Right == nil {
		return false
	}

	return (IsPerfectBinaryTree(node.Left, depth, level + 1) && IsPerfectBinaryTree(node.Right, depth, level + 1))
}

/**
Complete Binary Tree

A complete binary tree is a binary tree in which all the levels are completely filled except possibly the lowest one, which is filled from the left.
**/

func CountNodes(node *TreeNode) int{
	if node == nil {
		return 0
	}
	return 1 + CountNodes(node.Left) + CountNodes(node.Right)
}

func IsCompleteBinaryTree(node *TreeNode, index int, numberOfNodes int) bool {

	if node == nil {
		return true
	}

	if index >= numberOfNodes {
		return false
	}

	return (IsCompleteBinaryTree(node.Left, (2 * index + 1), numberOfNodes) && IsPerfectBinaryTree(node.Right, (2 * index + 2), numberOfNodes))
}



func main() {
	/**
	// Full Binary Tree
	tree := &Tree{}
	root:= &TreeNode{Val: 1}
	node1:= &TreeNode{Val: 2}
	node2:= &TreeNode{Val: 3}
	node3:= &TreeNode{Val: 4}
	node4:= &TreeNode{Val: 5}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4
	tree.Root = root
	
	fmt.Println(IsFullBinaryTree(tree.Root))
	**/

	/**
	// Perfect Binary Tree
	tree := &Tree{}
	root:= &TreeNode{Val: 1}
	node1:= &TreeNode{Val: 2}
	node2:= &TreeNode{Val: 3}
	node3:= &TreeNode{Val: 4}
	node4:= &TreeNode{Val: 5}
	node5:= &TreeNode{Val: 6}
	node6:= &TreeNode{Val: 7}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4
	root.Right.Left = node5
	root.Right.Right = node6

	tree.Root = root
	
	depth := CalculateTreeDepth(tree.Root)
	fmt.Println(fmt.Sprintf("Depth= %d", depth))
	fmt.Println(fmt.Sprintf("Is Perfect Tree: %t", IsPerfectBinaryTree(tree.Root, depth, 0)))
	**/

	/**
	// Complete Binary Tree
	tree := &Tree{}
	root:= &TreeNode{Val: 1}
	node1:= &TreeNode{Val: 2}
	node2:= &TreeNode{Val: 3}
	node3:= &TreeNode{Val: 4}
	node4:= &TreeNode{Val: 5}
	node5:= &TreeNode{Val: 6}
	node6:= &TreeNode{Val: 7}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4
	root.Right.Left = node5
	root.Right.Right = node6

	tree.Root = root
	
	count := CountNodes(tree.Root)
	fmt.Println(fmt.Sprintf("Count Nodes= %d", count))
	fmt.Println(fmt.Sprintf("Is Complete Binary tree Tree: %t", IsPerfectBinaryTree(tree.Root, 0, count)))
	**/

	

}