package mergeLists

// 功能: 合并K个升序链表
// 思路: 每个链表保存一个当前指针, 每次比较各个链表当前元素的大小, 取最小的, 然后移动到下一个

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	listArray := make([]*ListNode, len(lists))

	for lix, list := range lists {
		listArray[lix] = list
	}

	var head, curr *ListNode
	for {

		var index *int
		for lix, curr := range listArray {
			if curr == nil {
				continue
			}

			if index == nil {
				index = new(int)
				*index = lix
			} else {
				if curr.Val < listArray[*index].Val {
					*index = lix
				}
			}
		}

		if index == nil {
			break
		}

		if curr == nil {
			curr = listArray[*index]
			listArray[*index] = listArray[*index].Next
			head = curr
		} else {
			curr.Next = listArray[*index]
			curr = curr.Next
			listArray[*index] = listArray[*index].Next
		}
	}

	return head
}
