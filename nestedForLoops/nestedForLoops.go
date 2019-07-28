package nestedforloops

import (
	"fmt"
)

var arr = []int{1, 5, 10, 25}
var combinations []int
var originalCount int
var combinationsArray [][]int

// NestLoopsWrap wraps the nestloop function to initialize the combinations slice
func NestLoopsWrap(count int) {
	originalCount = count
	combinations = make([]int, count)
	NestLoops(count)
	fmt.Println(combinationsArray)
}

// NestLoops makes as many nested for loops as count and prints all combinations of elements in arr
func NestLoops(count int) {
	if count == 0 {
		fmt.Println(combinations)
		dst := make([]int, len(combinations))
		copy(dst, combinations)
		combinationsArray = append(combinationsArray, dst)
	} else {
		for i := range arr {
			combinations[originalCount-count] = arr[i]
			NestLoops(count - 1)
		}
	}
}
