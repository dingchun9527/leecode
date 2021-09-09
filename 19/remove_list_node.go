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

// 思路: 使用快慢指针的方法来获取位置
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	first, second := head, head
	var prev *ListNode

	for i := 0; i < n && first != nil; first, i = first.Next, i+1 {
	}

	for ; first != nil; first, second = first.Next, second.Next {
		prev = second
	}

	// 删除的是第一个节点
	if second == head {
		head = second.Next
	} else {
		prev.Next = second.Next
	}

	return head
}
