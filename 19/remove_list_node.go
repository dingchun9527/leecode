package removeListNode

// 功能: 剔除链表的第n个节点
// 思路:	遍历每个节点, 记录衣蛾hash map保存节点的地址, 然后再算倒数的

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	curr := head
	nodeMap := make(map[int]*ListNode)
	count := 0
	for ; curr != nil; curr, count = curr.Next, count+1 {
		nodeMap[count] = curr
	}

	offset := count - n
	node := nodeMap[offset]
	// 删除第一个
	if offset == 0 {
		if node.Next == nil {
			head = nil
		} else {
			head = node.Next
		}
	} else {
		prev := nodeMap[offset-1]
		if node.Next == nil {
			prev.Next = nil
		} else {
			prev.Next = node.Next
		}
	}

	return head
}
