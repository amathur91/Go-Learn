package main

import "fmt"
/**
Leetcode : First Missing Positive Number
Time Complexity : O(n)
Space Complexity : O(n)
Runtime: 4 ms, faster than 15.91% of Go online submissions for First Missing Positive.
Memory Usage: 2.4 MB, less than 6.73% of Go online submissions for First Missing Positive.
https://leetcode.com/problems/first-missing-positive/
 */
func main() {
	nums := []int{0}
	result := firstMissingPositive(nums)
	fmt.Println(result)
}

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	result := -1
	maxNum := findLargesNumInArray(nums)
	numMap := make(map[int]bool)
	populateMap(&numMap, &nums)
	for number := 1; number < maxNum; number++ {
		if !numMap[number] {
			result = number
			break
		}
	}
	if result == -1 {
		return maxNum + 1
	}
	return result
}

func populateMap(numMap *map[int]bool, nums *[]int){
	for index := 0; index < len(*nums); index++ {
		number := (*nums)[index]
		(*numMap)[number] = true
	}
}

func findLargesNumInArray(nums []int) int {
	maxNum := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] > maxNum {
			maxNum = nums[index]
		}
	}
	return maxNum
}