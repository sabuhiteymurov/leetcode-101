package main

import (
	"fmt"
	"leetcode-101/add-two-numbers"
	jumpGame "leetcode-101/jump-game"
	longestSubstringWithoutRepeatingChars "leetcode-101/longest-substring-without-repeating-chars"
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

	// Longest substring without repeating chars
	fmt.Println("3. Longest substring without repeating characters:", longestSubstringWithoutRepeatingChars.Solution("pwwkew"))

	// Jump Game
	fmt.Println("55. Jump Game:", jumpGame.CanJump([]int{2, 3, 1, 1, 4}))

}
