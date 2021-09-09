package reverseKGroup

// 功能: K 个一组翻转链表
// 思路: 使用一个长度为K的array保存链表节点, 逐个翻转着部分

type ListNode struct {
	Val  int
	Next *ListNode
}


func reverseKGroup(head *ListNode, k int) *ListNode {
	nodeArray := make([]*ListNode, k)
	for curr, i := head, 0; curr !=nil; curr, i=curr.Next,i+1 {
		// 达到K个, 将K个节点进行反转
		if i == k-1 {
			for j := k-1; j >0; j-- {
				nodeArray[j].Next = nodeArray[j-1]
			}
			nodeArray[0].Next = curr.Next
			i = 0
			continue
		}
		nodeArray[i] = curr
	}
	return head
}
