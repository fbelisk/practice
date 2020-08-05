package dp

import "fmt"

//输入: 3
//输出:
//[
//  [1,null,3,2],
//  [3,2,null,1],
//  [3,1,null,null,2],
//  [2,1,3],
//  [1,null,2,null,3]
//]
//解释:
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func generateTrees(n int) []*TreeNode {
//
//}
//
//func dfs(pre *TreeNode, remain []int) (node *TreeNode){
//	if len(remain) == 0 {
//		return
//	}
//	node.Left = dfs()
//	node.Right =
//}

//96. 不同的二叉搜索树
//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
//示例:
//
//输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3

func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i<=n; i++ {
		for j:=1; j <= i; j++ {
			dp[i] += dp[j - 1] * dp[i - j]
		}
	}
	fmt.Printf("%v", dp)
	return dp[n]
}
