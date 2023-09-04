package main

import (
	"math"
	"strconv"
)

type SpeSkillTest struct{}

func (r *SpeSkillTest) narcissistic(number int) bool {
	numberString := strconv.Itoa(number)
	numberDigits := len(numberString)
	sum := 0
	for _, digitStr := range numberString {
		digit := int(digitStr - '0')
		sum += int(math.Pow(float64(digit), float64(numberDigits)))
	}
	return sum == number
}

func (r *SpeSkillTest) findOutlier(arr []int) interface{} {
	var evenCount, oddCount int
	var even, odd int
	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
			even = num
		} else {
			oddCount++
			odd = num
		}
		if evenCount > 1 && oddCount == 1 {
			return odd
		} else if oddCount > 1 && evenCount == 1 {
			return even
		}
	}
	return false
}

func (r *SpeSkillTest) findNeedle(stack []string, needle string) int {
	for i, str := range stack {
		if str == needle {
			return i
		}
	}
	return -1
}

func (r *SpeSkillTest) blueOcean(arr1 []int, arr2 []int) []int {
	indexArr := make(map[int]bool)
	for _, num := range arr2 {
		indexArr[num] = true
	}
	var result []int
	for _, num := range arr1 {
		if !indexArr[num] {
			result = append(result, num)
		}
	}
	return result
}
