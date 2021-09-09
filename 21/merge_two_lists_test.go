package mergeTwoLists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list := MergeTwoLists(&ListNode{Val: 1}, &ListNode{Val: 2})
	assert.Equal(t, 1, list.Val)
}
