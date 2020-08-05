package bitOperation

//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
// 
//
//示例 1：
//
//输入：nums = [4,1,4,6]
//输出：[1,6] 或 [6,1]
//示例 2：
//
//输入：nums = [1,2,10,4,1,4,3,3]
//输出：[2,10] 或 [10,2]
// 
//
//限制：
//
//2 <= nums <= 10000
func singleNumbers(nums []int) []int {
	var res int
	for _, v := range nums {
		res ^= v
	}

	group := []int{0, 0}
	xor := res & -res
	for _, v := range nums {
		var index int
		//获取最低位1，对数据分组，出现两次的数字各自都会被分到同一组，出现一次的数字会被分到不同的组，各自异或就是结果
		if (xor & v) > 0 {
			index = 1
		}
		group[index] ^= v
	}

	return group
}

