package twoSum

func Solution(nums []int, target int) []int {
	seen := map[int]int{}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		j, existsInSeen := seen[diff]
		if existsInSeen {
			return []int{j, i}
		} else {
			seen[nums[i]] = i
		}
	}
	return []int{0, 0}
}
