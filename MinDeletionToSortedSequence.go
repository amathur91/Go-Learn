package main

import (
	"fmt"
)

func minDeletionForSortedSequence(arr *[]int) int {
	longestSequence := 1
	minNumArr := []int{}
	for index := 0; index < len(*arr); index++ {
		minNumArr = append(minNumArr, 1)
	}

	minNumArr[0] = 1
	for pivotIndex := 0; pivotIndex < len(*arr) ; pivotIndex++ {
		for compareIndex := 0; compareIndex < pivotIndex; compareIndex++ {
			if (*arr)[compareIndex] < (*arr)[pivotIndex] && minNumArr[pivotIndex] <= minNumArr[compareIndex] {
				minNumArr[pivotIndex] = minNumArr[compareIndex] + 1
			}
		}
	}

	for index := 0; index < len(minNumArr); index++ {
		if minNumArr[index] > longestSequence {
			longestSequence = minNumArr[index]
		}
	}
	return len(*arr) - longestSequence
}

func main() {
	arr := []int{4, 2, 3, 6, 10, 1, 12}
	var result int = minDeletionForSortedSequence(&arr)
	fmt.Printf("Minimum Deletions for Sorted Sequence : %d", result)
}
