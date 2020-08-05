package linkedList

func findDuplicate(nums []int) int {
	var a, b int
	a, b = 0, 0
	for {
		a = nums[a]
		b = nums[nums[b]]
		if a == b {
			break
		}
	}

	a = 0
	for {
		a = nums[a]
		b = nums[b]
		if a == b {
			return a
		}
	}
}
