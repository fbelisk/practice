package string

//解：双指针法
//344. 反转字符串
//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
//
//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//
//你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
func reverseString(s []byte)  {
	a := 0
	b := len(s) - 1
	if b < 0 {
		return
	}
	for {
		s[a], s[b] = s[b], s[a]
		a++
		b--
		if a >= b {
			return
		}
	}
}
