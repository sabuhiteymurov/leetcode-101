package jumpGame

func CanJump(nums []int) bool {
	n := len(nums)
	maxJump := 0

	if n == 1 {
		return true
	}

	for i := 0; i < n-1 && maxJump >= i; i++ {
		if (i + nums[i]) > maxJump {
			maxJump = i + nums[i]
		}

		if maxJump >= n-1 {
			return true
		}
	}

	return false
}
