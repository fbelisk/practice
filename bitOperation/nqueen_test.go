package bitOperation

import (
	"go.uber.org/zap"
	"struct/components"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	components.InitLogTest()
	r := solveNQueens(4)
	components.L.Debug("zap", zap.Any("r", r))
}

