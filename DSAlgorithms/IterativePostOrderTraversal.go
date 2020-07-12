package main

import "fmt"
/**
Very Optimized Solution for the Iterative Postorder Tree Traversal
Time Complexity : O(n)
Space Complexity : O(n)

Leetcode Result :
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2.1 MB, less than 15.96% of Go online submissions for Binary Tree Postorder Traversal.

Link :
https://leetcode.com/problems/binary-tree-postorder-traversal
Category : Hard
 */
func main() {
	rootNode := createNode(3)
	rootNode.Left = createNode(1)
	rootNode.Right = createNode(2)
	size := 0
	visitedMap := make(map[*TreeNode]bool)
	countTotalNodes(rootNode, &size, &visitedMap)
	stack := Stack{treeNodes: make([]*TreeNode, size), currentPointer: size -1}
	result := make([]int, size)

	index := 0
	pushToStack(rootNode, &stack)
	traverseTree(&stack, &size, &result, &index, &visitedMap)
	fmt.Println(result)
}

func traverseTree(stack *Stack, size *int, result *[]int, index *int, visitedMap *map[*TreeNode]bool) {
	for ;stack.currentPointer != *size - 1 ; {
		node := popFromStack(stack)
		if !(*visitedMap)[node]  {
			//explore it and push it back
			pushToStack(node, stack)

			if node.Right != nil {
				pushToStack(node.Right, stack)
			}

			if node.Left != nil {
				pushToStack(node.Left, stack)
			}

			(*visitedMap)[node] = true
		} else {
			(*result)[*index] = node.Val
			*index++
		}
	}
}

func pushToStack(node *TreeNode, stack *Stack) {
	stack.treeNodes[stack.currentPointer] = node
	stack.currentPointer--
}

func popFromStack(stack *Stack) *TreeNode {
	node := stack.treeNodes[stack.currentPointer + 1]
	stack.currentPointer++
	return node
}

func countTotalNodes(rootNode *TreeNode, size *int, visitedMap *map[*TreeNode]bool) {
	if rootNode != nil {
		*size++
		(*visitedMap)[rootNode] = false
		countTotalNodes(rootNode.Left, size, visitedMap)
		countTotalNodes(rootNode.Right, size, visitedMap)
	}
}


func createNode(dataValue int) *TreeNode {
	node := TreeNode{Val: dataValue, Left: nil, Right: nil}
	return &node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	treeNodes []*TreeNode
	currentPointer int
}
