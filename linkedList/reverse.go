package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表题解1
func ReverseList1(head *ListNode) *ListNode {
	if head.Next == nil {
		head.Next = head
		return head
	}
	nextNode := head.Next
	head.Next = nil
	first := ReverseList1(nextNode)
	nextNode.Next = head
	return first
}

//反转链表题解2
func ReverseList2(head *ListNode) *ListNode {
	var prev, cur *ListNode
	prev, cur = nil, head
	for {
		cur.Next, prev, cur = prev, cur, cur.Next
		if cur == nil {
			break
		}
	}
	return prev
}

//题目
//这其实是一道变形的链表反转题，大致描述如下 给定一个单链表的头节点 head,实现一个调整单链表的函数，使得每K个节点之间为一组进行逆序，并且从链表的尾部开始组起，头部剩余节点数量不够一组的不需要逆序。（不能使用队列或者栈作为辅助）
//
//例如：
//
//链表:1->2->3->4->5->6->7->8->null, K = 3。那么 6->7->8，3->4->5，1->2各位一组。调整后：1->2->5->4->3->8->7->6->null。其中 1，2不调整，因为不够一组。

func ReverseList3(head *ListNode) *ListNode {
	var prev, cur *ListNode
	prev, cur = nil, head
	for {
		cur.Next, prev, cur = prev, cur, cur.Next
		if cur == nil {
			break
		}
	}
	return prev
}
