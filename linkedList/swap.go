package linkedList

//链表元素两两交换
func swapPairs(head *ListNode) *ListNode {
	var prev, cur, next, newHead *ListNode
	newHead, prev, cur, next  = head.Next, nil, head, head.Next
	for {
		if prev != nil {
			prev.Next = next
		}
		cur.Next, next.Next, prev = next.Next, cur, cur
		if cur.Next == nil || cur.Next.Next == nil {
			break
		}
		cur, next = cur.Next, cur.Next.Next
	}
	return newHead
}
