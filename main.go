package main

import "log"

func main() {
	speSkillTest := SpeSkillTest{}

	// challenge 1
	log.Println("narcissistic(153)", speSkillTest.narcissistic(153))
	log.Println("narcissistic(111)", speSkillTest.narcissistic(111))

	// challenge 2
	log.Println("[2, 4, 0, 100, 4, 11, 2602, 36]:", speSkillTest.findOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	log.Println("[160, 3, 1719, 19, 11, 13, -21]:", speSkillTest.findOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	log.Println("[11, 13, 15, 19, 9, 13, -21]:", speSkillTest.findOutlier([]int{11, 13, 15, 19, 9, 13, -21}))

	// challenge 3
	log.Println(`findNeedle(["red", "blue", "yellow", "black", "grey"], "blue"):`, speSkillTest.findNeedle([]string{"red", "blue", "yellow", "black", "grey"}, "blue"))

	// challenge 4
	log.Println("blueOcean([1,2,3,4,6,10], [1]):", speSkillTest.blueOcean([]int{1, 2, 3, 4, 6, 10}, []int{1}))
	log.Println("blueOcean([1,5,5,5,5,3], [5]):", speSkillTest.blueOcean([]int{1, 5, 5, 5, 5, 3}, []int{5}))
}
