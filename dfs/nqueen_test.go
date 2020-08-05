package dfs

import (
	"go.uber.org/zap"
	"struct/components"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	components.InitLogTest()
	r := solveNQueens(16)
	components.L.Debug("zap", zap.Any("r", r))
}

