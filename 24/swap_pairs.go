package swapPairs

// 功能: 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 思路: 第一个节点指向第二个节点的下一个节点, 第二个节点指向第一个节点, 以此类推

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	var prev, curr, next *ListNode
	for curr = head; curr != nil && curr.Next != nil; curr = curr.Next {
		next = curr.Next
		curr.Next = next.Next
		next.Next = curr

		if prev == nil {
			prev = curr
			head = next
		} else {
			prev.Next = next
			prev = curr

		}
	}

	return head
}
