package main

import (
	"fmt"
	"leetcode-101/two-sum"
)

func main() {
	twoSumOutput := twoSum.Solution([]int{1, 2, 5, 4, 7}, 7)
	fmt.Println("Two sum solution output:", twoSumOutput[0], twoSumOutput[1])
}
