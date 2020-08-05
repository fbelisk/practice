package quickSort

type QuickSort struct {
	Array []int
}

//取第一个元素为基准块，j先从右向左收缩，遇到比基准块小的就互相交换，然后i从左向右收缩，遇到比基准块打的就交换，最后i和j会相等，落到基准块的位置
func (q *QuickSort) Sort(beginIndex, endIndex int) {
	if endIndex <= beginIndex {
		return
	}
	baseValue := q.Array[beginIndex]
	i := beginIndex
	j := endIndex
	for j > i {
		for j > i && q.Array[j] > baseValue {
			j--
		}
		if j > i {
			//交换
			q.Array[i] = q.Array[j]
			i++
		}
		for i < j && q.Array[i] < baseValue {
			i++
		}
		if i < j {
			//交换
			q.Array[j] = q.Array[i]
			j--
		}
	}
	//不需要把基准元素填充到对应位置，因为基准元素的位置下一轮还会被其他元素块填充，所以最后填充即可
	q.Array[j] = baseValue
	q.Sort(beginIndex, j -1)
	q.Sort(j+1, endIndex)
}