package dp

//53. 最大子序和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//dp[i] 定义 包含当前元素的最大子序和

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for i:=1;i<len(nums);i++ {
		dp[i] = max(dp[i - 1] + nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}