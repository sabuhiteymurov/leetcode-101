package main

import (
	"fmt"
	"leetcode-101/add-two-numbers"
	"leetcode-101/two-sum"
)

func main() {
	// Two sum
	twoSumOutput := twoSum.Solution([]int{1, 2, 5, 4, 7}, 7)
	fmt.Println("1. Two sum solution output:", twoSumOutput[0], twoSumOutput[1])

	// Add two numbers
	l1 := &addTwoNumbers.ListNode{2, &addTwoNumbers.ListNode{4, &addTwoNumbers.ListNode{3, nil}}}
	l2 := &addTwoNumbers.ListNode{5, &addTwoNumbers.ListNode{6, &addTwoNumbers.ListNode{4, nil}}}
	addTwoNumbersOutput := addTwoNumbers.Solution(l1, l2)
	fmt.Println("2. Add two numbers solution output:", addTwoNumbersOutput)
}
