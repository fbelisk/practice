package linkedList

import (
	"go.uber.org/zap"
	"rb-tree/components"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	components.InitLogTest()

	d := &ListNode{
		Val:4,
		Next: nil,
	}
	c := &ListNode{
		Val:3,
		Next: d,
	}
	b := &ListNode{
		Val:2,
		Next: c,
	}
	a := &ListNode{
		Val:1,
		Next: b,
	}
	head := swapPairs(a)
	components.L.Debug("zap", zap.Any("head", head))
}
