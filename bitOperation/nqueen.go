package bitOperation

import (
	"math"
)

type Queen struct {
	n      int
	Tmp    []int
	Result [][]int
}

func solveNQueens(n int) [][]string {
	q := Queen{
		n:      n,
		Tmp:    []int{},
		Result: [][]int{},
	}
	q.dfs(0, 0, 0, 0)
	return generateResult(q.Result, q.n)
}

func (q *Queen) dfs(row, col, pie, na int) {
	//递归终止条件
	if row >= q.n {
		//一定要进行内存拷贝，不然追加进的都是同一个tmp，最后输出的结果都是相同的slice元素
		c := make([]int, q.n)
		copy(c, q.Tmp)
		//保存结果集
		q.Result = append(q.Result, c)
		return
	}
	//获取可用单位空位，所有排除条件进行或运算 ，再取反就可获取所有允许的位为1的值， & (1<<uint(q.n) - 1)为清除高位的1
	bits := ^(col | pie | na) & (1<<uint(q.n) - 1)
	for {
		if bits == 0 {
			break
		}
		//每次获取最低位的1作为本次的皇后放置位置
		current := bits & -bits
		//放入临时结果集
		q.Tmp = append(q.Tmp, current)
		//下一次递归，注意撇和捺要进行左右移操作
		q.dfs(row+1, col|current, (pie|current)<<1, (na|current)>>1)
		//递归返回时清空递归产生的元素
		q.Tmp = q.Tmp[0:len(q.Tmp) - 1]
		//清空最低位的1
		bits = bits & (bits - 1)
	}
}

//格式化输出格式
func generateResult(res [][]int, n int) (result [][]string) {
	for _, v := range res {
		var s []string
		for _, val := range v {
			p := n - int(math.Log(float64(val))) - 2
			str := ""
			for i := 0; i < n; i++ {
				if i == p {
					str += "Q"
				} else {
					str += "."
				}
			}
			s = append(s, str)
		}
		result = append(result, s)
	}
	return
}
