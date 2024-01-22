package jumpGame

func CanJump(nums []int) bool {
	var canJump bool

	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums) && len(nums) != 1 {
			break
		}
		if nums[i] == 0 {
			if len(nums) == 1 {
				canJump = true
			}
			continue
		} else if (nums[i] + i) >= len(nums)-1 {
			canJump = true
		}

	}

	return canJump
}
