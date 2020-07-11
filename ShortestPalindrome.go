package main

import "fmt"

func main() {
	input := "abb"
	finalPalindrome := shortestPalindrome(input)
	fmt.Printf("The palindrome formed is : %s ", finalPalindrome)
}

func shortestPalindrome(input string) string {
	palindromeSubsequenceMatrix := make([][]int, len(input))
	for row := range palindromeSubsequenceMatrix {
		palindromeSubsequenceMatrix[row] = make([]int, len(input))
	}
	buildLongestPalindromicSubsequence(&input, &palindromeSubsequenceMatrix)
	return buildShortestPalindrome(&input, &palindromeSubsequenceMatrix)
}

func buildShortestPalindrome(input *string, palindromeSubsequenceMatrix *[][]int) string{
	start := 0
	workingString := *input
	end := len(*input) - 1
	runes := []rune(*input)
	checkAndCompute(&start, &end, &workingString, &runes, palindromeSubsequenceMatrix)
	return workingString
}

func checkAndCompute(startIndex *int, endIndex *int, workingString *string, input *[]rune, matrix *[][]int) {
	if *startIndex <= *endIndex {
		leftCharacter := string((*input)[*startIndex])
		rightCharacter := string((*input)[*endIndex])
		if leftCharacter == rightCharacter {
			*startIndex++
			*endIndex--
			checkAndCompute(startIndex, endIndex, workingString, input, matrix)
		}else{
			leftLPS := (*matrix)[*startIndex][*endIndex - 1]
			rightLPS := (*matrix)[*startIndex + 1][*endIndex]
			if leftLPS == rightLPS {
				leftPart := string((*input)[:*startIndex])
				middlePart := string((*input)[*startIndex:*endIndex])
				rightPart := string((*input)[*endIndex:])
				*workingString = leftPart + rightCharacter + middlePart + rightPart
			}else if leftLPS > rightLPS {
				leftPart := string((*input)[:*startIndex])
				middlePart := string((*input)[*startIndex:*endIndex+1])
				rightPart := string((*input)[*endIndex+1:])
				*workingString = leftPart + rightCharacter + middlePart + rightPart
			}else if leftLPS < rightLPS {
				leftPart := string((*input)[:*startIndex])
				middlePart := string((*input)[*startIndex:*endIndex+1])
				rightPart := string((*input)[*endIndex+1:])
				*workingString = leftPart  + middlePart + leftCharacter + rightPart
			}
			*endIndex = len(*workingString) - 1
			*startIndex = 0
			*input = []rune(*workingString)
			checkAndCompute(startIndex, endIndex , workingString, input, matrix)
		}
	}
}

func buildLongestPalindromicSubsequence(input *string, palindromeSubsequenceMatrix *[][]int) {
	runes := []rune(*input)
	for windowSize := 0; windowSize < len(*input); windowSize++ {
		startWindow := 0
		for endWindow := startWindow + windowSize; endWindow < len(*input); endWindow++ {
			if endWindow-startWindow == 0 {
				(*palindromeSubsequenceMatrix)[startWindow][endWindow] = 1
			} else {
				leftCharacter := string(runes[startWindow])
				rightCharacter := string(runes[endWindow])
				if leftCharacter == rightCharacter {
					(*palindromeSubsequenceMatrix)[startWindow][endWindow] = 2 + (*palindromeSubsequenceMatrix)[startWindow+1][endWindow-1]
				} else {
					(*palindromeSubsequenceMatrix)[startWindow][endWindow] = getMax((*palindromeSubsequenceMatrix)[startWindow+1][endWindow],
						(*palindromeSubsequenceMatrix)[startWindow][endWindow-1])
				}
			}
			startWindow++
		}
	}
}

func getMax(value1 int, value2 int) int{
	if value1 > value2 {
		return value1
	}
	return value2
}