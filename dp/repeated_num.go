package dp

//给定正整数 N，返回小于等于 N 且具有至少 1 位重复数字的正整数的个数。
func numDupDigitsAtMostN(N int) int {
	dp := make(map[int]int)
	for i := 1; i <= N; i++ {
		if hasRepeated(i) {
			dp[i] = dp[i - 1] + 1
		} else {
			dp[i] = dp[i - 1]
		}
	}
	return dp[N]
}

func hasRepeated(N int) bool{
	return true
}

func containsDuplicate(nums []int) bool {
	var dups uint64
	for _, n := range nums {
		exists := dups & (1 << uint64(n))
		if exists == 0 {
			dups |= 1 << uint64(n)
		} else {
			return true
		}
	}
	return false
}