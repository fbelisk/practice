package dp

import "fmt"

//dp[i] = dp[i - coin] + 1
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	dp := make(map[int]int, amount + 1)
	dp[0] = 0
	for i:= 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	fmt.Printf("%v", dp)
	return dp[amount]
}
