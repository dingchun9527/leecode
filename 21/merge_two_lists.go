package mergeTwoLists

// 功能: 合并两个有序链表
// 思路: 遍历两个链表, 取当前结点为更小的

type ListNode struct {
	Val  int
	Next *ListNode
}


func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, curr *ListNode

	for curr1, curr2:= l1, l2; curr1 != nil || curr2 !=nil; {

		if curr1 != nil && curr2 != nil {
			if curr1.Val < curr2.Val {
				if curr == nil {
					curr = curr1
					head = curr
				} else {
					curr.Next = curr1
					curr = curr.Next
				}
				curr1 = curr1.Next
			} else {
				if curr == nil {
					curr = curr2
					head = curr
				} else {
					curr.Next = curr2
					curr = curr.Next
				}
				curr2 = curr2.Next
			}

		} else if curr1 != nil {
			if curr == nil {
				curr = curr1
				head = curr
			} else {
				curr.Next = curr1
				curr = curr.Next
			}
			curr1 = curr1.Next
		} else {
			if curr == nil {
				curr = curr2
				head = curr
			} else {
				curr.Next = curr2
				curr = curr.Next
			}
			curr2 = curr2.Next
		}
	}

	return head
}