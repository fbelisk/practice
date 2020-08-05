package linkedList

import (
	"go.uber.org/zap"
	"rb-tree/components"
	"testing"
)

func TestReverseList(t *testing.T) {
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
	head := ReverseList2(a)
	components.InitLogTest()
	components.L.Debug("zap", zap.Any("head", head))
}
