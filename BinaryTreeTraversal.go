package main

import (
	"fmt"
	"time"
)

type Node struct {
	data int32
	leftChild *Node
	rightChild *Node
}

func newNode(dataValue int32) Node {
	node := Node{data: dataValue, leftChild: nil, rightChild: nil}
	return node
}

func inorderTraversal(rootNode *Node) {
	if rootNode != nil {
		inorderTraversal(rootNode.leftChild)
		fmt.Printf(" %d ", rootNode.data)
		inorderTraversal(rootNode.rightChild)
	}
}

func getNewNodeAddress(dataValue int32) *Node{
	tempNode := newNode(dataValue)
	return &tempNode
}

func main() {
	rootNode := newNode(1)
	rootNode.leftChild = getNewNodeAddress(2)
	rootNode.rightChild = getNewNodeAddress(3)
	rootNode.leftChild.leftChild = getNewNodeAddress(4)
	rootNode.leftChild.rightChild = getNewNodeAddress(5)
	rootNode.rightChild.leftChild = getNewNodeAddress(6)
	rootNode.rightChild.rightChild = getNewNodeAddress(7)
	rootNode.leftChild.leftChild.leftChild = getNewNodeAddress(8)
	var height int32 = 1
	start := time.Now()
	findHeightOfTree(&rootNode, &height, 0)

	fmt.Printf("Maximum Height Of the Tree : %d. \n", height)
	fmt.Printf("Overall Execution Time : %v \n", (time.Since(start).Microseconds()))
	//inorderTraversal(&rootNode)
}

func findHeightOfTree(rootNode *Node, height *int32, currentHeight int32) {
	if rootNode != nil{
		if currentHeight > *height {
			*height = currentHeight
		}
		findHeightOfTree(rootNode.leftChild, height, currentHeight + 1)
		findHeightOfTree(rootNode.rightChild, height, currentHeight + 1)
	}
}
