package dfs

//322. 零钱兑换
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

type Coin struct {
	mini  uint
	coins []int
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	c := &Coin{
		mini: 0,
		coins: coins,
	}
	c.dfs(amount, 0)
	if c.mini > 0 {
		return int(c.mini)
	}
	return -1
}

func (c *Coin) dfs(remain int, current uint) {
	if remain == 0 {
		if current < c.mini || c.mini == 0{
			c.mini = current
		}
		return
	}

	for i := 0; i < len(c.coins); i++ {
		if remain - c.coins[i] < 0 {
			continue
		}

		c.dfs(remain - c.coins[i], current + 1)
	}
}
