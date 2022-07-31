package solution

// dp

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i; j <= i+nums[i]; j++ {
			if j >= len(nums)-1 || dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}
