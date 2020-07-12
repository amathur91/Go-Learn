package main

import "fmt"

func main() {
	s1 := "distance"
	s2 := "springbok"
	result := findMinDistance(&s1, &s2)
	fmt.Printf("Minimum Distance is %d", result)
}

func findMinDistance(s1 *string, s2 *string) int {
	if len(*s1) == 0 {
		return len(*s2)
	}

	if len(*s2) == 0 {
		return len(*s1)
	}

	matrix := make([][]int, len(*s1)+1)
	for row := range matrix {
		matrix[row] = make([]int, len(*s2)+1)
	}

	for s1Index := 0; s1Index <= len(*s1); s1Index++ {
		matrix[s1Index][0] = s1Index
	}
	for s2Index := 0; s2Index <= len(*s2); s2Index++ {
		matrix[0][s2Index] = s2Index
	}

	s1Chars := []rune(*s1)
	s2Chars := []rune(*s2)

	for s1Index := 1; s1Index <= len(*s1); s1Index++ {
		for s2Index := 1; s2Index <= len(*s2); s2Index++ {
			s1Character := string(s1Chars[s1Index - 1])
			s2Character := string(s2Chars[s2Index - 1])

			if s1Character == s2Character {
				matrix[s1Index][s2Index] = matrix[s1Index-1][s2Index-1]
			} else {
				matrix[s1Index][s2Index] = 1 + getMinValue(
					getMinValue(matrix[s1Index-1][s2Index],   //remove
					matrix[s1Index - 1][s2Index - 1]), //replace
					matrix[s1Index][s2Index-1]) //insert
			}
		}
	}

	return matrix[len(*s1)][len(*s2)]
}


func getMinValue(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}