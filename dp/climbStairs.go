package dp

//70. 爬楼梯
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	count := make([]int, n+1)
	count[1] = 1
	count[2] = 2
	for i := 3; i <= n; i++ {
		count[i] = count[i-2] + count[i-1]
	}
	return count[n]
}
