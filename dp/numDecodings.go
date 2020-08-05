package dp

//91. 解码方法
//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
//示例 1:
//
//输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//示例 2:
//
//输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	zero := "0"[0]
	max := "2"[0]<<8 + "6"[0]
	for k:=0; k < l; k++{
		if k == 0 {
			if s[k] == zero{
				dp[k] = 0
			} else {
				dp[k] = 1
			}
			continue
		}

		if s[k] != zero && s[k-1] == zero && dp[k] == 0{
			dp[k] = 1
		}

		if s[k] != zero && s[k-1] != zero && s[k - 1] << 8 + s[k] <= max{
			dp[k] = dp[k - 1] + 1
			continue
		}

		dp[k] = dp[k - 1]
	}
	return dp[l - 1]
}
