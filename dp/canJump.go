package dp

//dp[i] 总长为i的数组，最大位置下标
func canJump(nums []int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	dp := make([]int, l + 1)
	dp[1] = nums[0]
	for i := 2; i <= l; i++{
		if dp[i - 1] >= i - 1 {
			dp[i] = max(dp[i - 1], i - 1 + nums[i - 1])
			continue
		}
		dp[i] = dp[i - 1]
	}
	return dp[l - 1] >= len(nums) - 1
}
