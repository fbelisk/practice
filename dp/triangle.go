package dp

//120. 三角形最小路径和
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//例如，给定三角形：
//
//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//说明：
//
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	min := make([][]int, l)
	for i := l - 1; i >= 0; i-- {
		min[i] = make([]int, len(triangle[i]))
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if i == l - 1 {
				min[i][j] = triangle[i][j]
				continue
			}
			if min[i+1][j] < min[i+1][j + 1] {
				min[i][j] = min[i+1][j] + triangle[i][j]
			} else {
				min[i][j] = min[i+1][j + 1] + triangle[i][j]
			}
		}
	}
	return min[0][0]
}
