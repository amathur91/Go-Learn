package main

import (
	"fmt"
)

/**
Leetcode Problem : https://leetcode.com/problems/validate-binary-search-tree/
Level : Medium
Time Complexity: O(n)

Results :
Runtime: 4 ms, faster than 98.38% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.4 MB, less than 87.05% of Go online submissions for Validate Binary Search Tree.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func main() {
	root := TreeNode{Val: -2147483648}
	//root.Left = &TreeNode{Val: }
	root.Right = &TreeNode{Val: 2147483647}
	result := isValidBST(&root)
	fmt.Println("The result of the input BST is", result)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return validateNode(root, -2147483649, 2147483648)
}

func validateNode(root *TreeNode, minBound int, maxBound int) (bool) {
	if root != nil {
		if root.Val > minBound && root.Val < maxBound {
			//we are good with the current node. recursively call for the left and right child
			return validateNode(root.Left, minBound, root.Val) && validateNode(root.Right, root.Val, maxBound)
		} else {
			//the current node does not fall between the range and hence is not valid
			return false
		}
	}
	return true
}