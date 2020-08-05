package hash

import "sort"

func TwoSum(nums[]int, target int) []int {
	numMap := make(map[int]int)
	if len(nums) == 0 {
		return nil
	}
	for i, v := range nums {
		match := target - v
		index, ok := numMap[match]
		if ok {
			return []int{index, i}
		}
		numMap[v] = i
	}
	return nil
}

//
func ThreeSum(nums[]int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return nil
	}
	//排序、方便去重
	sort.Ints(nums)
	for i1, v1 := range nums {
		//和前一个数字比较，相同的数字只遍历一遍
		if i1 > 0 && v1 == nums[i1 - 1] {
			continue
		}
		target := - v1
		numMap := make(map[int]int)
		//只取nums[i1+1:]，向后遍历，以防重复遍历
		for _, v2 := range nums[i1+1:] {
			_, ok := numMap[v2]
			if ok {
				res = append(res, []int{v1, v2, target - v2})
			} else {
				numMap[target - v2] = 1
			}
		}
	}
	return res
}
