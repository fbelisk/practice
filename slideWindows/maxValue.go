package slideWindows

//滑动窗口最大值
//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
func MaxSlidingWindow1(nums []int, k int) []int {
	var pStart, pEnd, pMax int
	var maxs []int
	pEnd = k - 1
	numslen := len(nums)
	if numslen == 0 || k == 0 {
		return nums
	}
	for {
		if pEnd >= numslen {
			break
		}

		if pMax < pStart {
			pMax = pStart
		}

		offset := pMax+1
		for i, v := range nums[pMax+1 : pEnd+1] {
			if nums[pMax] <= v {
				pMax = offset + i
			}
		}
		println(pMax, nums[pMax])
		maxs = append(maxs, nums[pMax])
		pEnd++
		pStart++
	}
	return maxs
}

func MaxSlidingWindow2(nums []int, k int) []int {
	if k == 0 {
		return nil
	}
	if k ==1 || len(nums) == 0 {
		return nums
	}
	windows := make([]int, 0)
	res := make([]int, 0)
	for i, v := range nums {
		//窗口移动 移除第一个元素
		if len(windows) > 0 && i >= k && windows[0] <= i - k{
			windows = windows[1:]
		}
		//窗口移动 塞入新元素,同时进行比较，把最大元素左边元素统统pop出去
		for {
			if len(windows) > 0 && v >= windows[len(windows) - 1] {
				windows = windows[:len(windows) - 1]
			} else {
				break
			}
		}
		windows = append(windows, i)
		//塞入结果集 (前几个元素不塞，因为前几次是在初始化滑动窗口)
		if i >= k - 1 {
			res = append(res, nums[windows[0]])
		}
	}
	return res
}