package main

import "fmt"

/**
Solution for the Edit Distance where we transform one string to another using
three operations like replace, remove, insert.
This is a recursive solution and the time complexity is O (3^n)

 */
func main() {
	s1 := "passpot"
	s2 := "ppsspqrt"
	result := minEditOperations(&s1, &s2)
	fmt.Printf("Minimum Edit Operations : %d", result)
}

func minEditOperations(s1 *string, s2 *string) int {
	s1Characters := []rune(*s1)
	s2Characters := []rune(*s2)
	return minOperationsUtils(s1, s2, 0, 0, &s1Characters, &s2Characters)
}

func minOperationsUtils(s1 *string, s2 *string, s1Index int, s2Index int, s1Characters *[]rune, s2Characters *[]rune) int {
	if s1Index < len(*s1) && s2Index < len(*s2) {
		//if we havent reach the end of any of the string
		s1Char := string((*s1Characters)[s1Index])
		s2Char := string((*s2Characters)[s2Index])
		if s1Char == s2Char {
			return minOperationsUtils(s1, s2, s1Index + 1, s2Index + 1, s1Characters, s2Characters)
		}else{
			//we need to apply any one of the operations here
			//replace
			result1 := 1 +  minOperationsUtils(s1, s2, s1Index + 1, s2Index + 1, s1Characters, s2Characters)
			//remove
			result2 := 1 + minOperationsUtils(s1, s2, s1Index + 1, s2Index, s1Characters, s2Characters)

			//insert
			result3 := 1 + minOperationsUtils(s1, s2, s1Index, s2Index + 1, s1Characters, s2Characters)

			return minValue(result1, result2, result3)
		}

	}else{
		//find out which one's length is exhausted and proceed accordingly.
		if s2Index < len(*s2) {
			return len(*s2) - s2Index
		}else if s1Index < len(*s1) {
			return len(*s1) - s1Index
		}else{
			return 0
		}
	}
}

func minValue(num1 int, num2 int, num3 int) int {
	if num1 < num2 && num1 < num3{
		return num1
	}

	if num2 < num1 && num2 < num3 {
		return num2
	}

	return num3
}
