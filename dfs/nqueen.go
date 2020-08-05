package dfs

func solveNQueens(n int) [][]string {
	result := &[][]int{}
	pies := make(map[int]bool)
	nas := make(map[int]bool)
	cols := make(map[int]bool)
	tmp := make([]int, 0)
	dfs(n, 0, result, pies, nas, cols, tmp)
	return generateResult(*result, n)
}

func dfs(n, row int, result *[][]int, pies, nas, cols map[int]bool, tmp []int) {
	if row >= n {
		(*result)=append((*result),tmp)
		return
	}

	for col := 0; col < n; col++ {
		if row == 0 {
			tmp = []int{}
		}
		v1, ok1 := cols[col]
		v2, ok2 := pies[col+row]
		v3, ok3 := nas[col-row]
		if (ok1 && v1) || (ok2 && v2) || (ok3 && v3) {
			continue
		}
		cols[col] = true
		pies[col+row] = true
		nas[col-row] = true
		tmp = append(tmp, col)
		dfs(n, row + 1, result, pies, nas, cols, tmp)
		cols[col] = false
		pies[col+row] = false
		nas[col-row] = false
	}
}

func generateResult(res [][]int, n int) (result [][]string) {
	for _, v := range res {
		var s []string
		for _, val := range v {
			str := ""
			for i := 0; i < n; i++ {
				if i == val {
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
