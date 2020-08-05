package bitOperation

func hammingWeight(num uint32) int {
	var count int
	if num == 0 {
		return count
	}
	for {
		if num = num & (num -1); num > 0 {
			count ++
		} else {
			return count + 1
		}
	}
}