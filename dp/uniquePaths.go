package dp

func uniquePaths(m int, n int) int {
	dp := make([][]int, m + 1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n + 1)
		for j := 1; j <= n; j++ {
			if j == 1 || i == 1{
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if j == 0 && i == 0{
				dp[i][j]  = 1
				continue
			}

			if j == 0 {
				dp[i][j] = dp[i - 1][j]
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j - 1]
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m - 1][n - 1]
}