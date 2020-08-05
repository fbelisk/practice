package dp

import (
	"go.uber.org/zap"
	"struct/components"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	components.InitLogTest()
	tr := [][]int{
		{2},{3,4},{6,5,7},{4,1,8,3},
	}
	r := minimumTotal(tr)
	components.L.Debug("zap", zap.Any("r", r))
}

