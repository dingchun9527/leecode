package sumOfTwoNumbers

import "fmt"

// 链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 功能: 两个链表的中的数字相加得到新的链表
// 思路: 遍历两个链表, 取两数相加(节点为null则是0), 注意进位
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	var l3, pos3 *ListNode

	for pos1, pos2 := l1, l2; pos1 != nil || pos2 != nil; {
		num1, num2 := 0, 0
		if pos1 != nil {
			num1 = pos1.Val
			pos1 = pos1.Next
		}

		if pos2 != nil {
			num2 = pos2.Val
			pos2 = pos2.Next
		}
		data := num1 + num2 + carry
		carry = data / 10
		if pos3 == nil {
			pos3 = &ListNode{Val: data % 10}
			l3 = pos3
		} else {
			pos3.Next = &ListNode{Val: data % 10}
			pos3 = pos3.Next
		}
	}

	if carry != 0 {
		pos3.Next = &ListNode{Val: carry}
		pos3 = pos3.Next
	}

	return l3
}

func (l *ListNode) InsertAtHead(data int) *ListNode {
	if l == nil {
		return &ListNode{Val: data}
	}
	tmp := l
	l = &ListNode{Val: data}
	l.Next = tmp

	return l
}
func (l *ListNode) Show() {

	for pos := l; pos != nil; pos = pos.Next {
		fmt.Printf("%d ", pos.Val)
	}
	fmt.Println()
}

func (l *ListNode) Equal(list *ListNode) bool {
	pos1, pos2 := l, list
	for ; pos1 != nil && pos2 != nil; pos1, pos2 = pos1.Next, pos2.Next {
		if pos1.Val != pos2.Val {
			return false
		}
	}

	if pos1 != nil || pos2 != nil {
		return false
	}

	return true
}
