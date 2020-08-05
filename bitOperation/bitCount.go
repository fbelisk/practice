package bitOperation

//338. 比特位计数
//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
func countBits(num int) []int {
	count := make([]int, num + 1)
	for i := 1; i <= num; i++ {
		count[i] = count[i&(i-1)] + 1
	}
	return count
}
